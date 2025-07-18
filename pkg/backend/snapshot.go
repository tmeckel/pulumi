// Copyright 2016-2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package backend

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"golang.org/x/exp/slices"

	"github.com/pulumi/pulumi/pkg/v3/engine"
	"github.com/pulumi/pulumi/pkg/v3/resource/deploy"
	"github.com/pulumi/pulumi/pkg/v3/secrets"
	"github.com/pulumi/pulumi/sdk/v3/go/common/env"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	"github.com/pulumi/pulumi/sdk/v3/go/common/version"
)

// DisableIntegrityChecking can be set to true to disable checkpoint state integrity verification.  This is not
// recommended, because it could mean proceeding even in the face of a corrupted checkpoint state file, but can
// be used as a last resort when a command absolutely must be run.
var DisableIntegrityChecking bool

// SnapshotPersister is an interface implemented by our backends that implements snapshot
// persistence. In order to fit into our current model, snapshot persisters have two functions:
// saving snapshots and invalidating already-persisted snapshots.
type SnapshotPersister interface {
	// Persists the given snapshot. Returns an error if the persistence failed.
	Save(snapshot *deploy.Snapshot) error
}

// SnapshotManager is an implementation of engine.SnapshotManager that inspects steps and performs
// mutations on the global snapshot object serially. This implementation maintains two bits of state: the "base"
// snapshot, which is completely immutable and represents the state of the world prior to the application
// of the current plan, and a "new" list of resources, which consists of the resources that were operated upon
// by the current plan.
//
// Important to note is that, although this SnapshotManager is designed to be easily convertible into a thread-safe
// implementation, the code as it is today is *not thread safe*. In particular, it is not legal for there to be
// more than one `SnapshotMutation` active at any point in time. This is because this SnapshotManager invalidates
// the last persisted snapshot in `BeginSnapshot`. This is designed to match existing behavior and will not
// be the state of things going forward.
//
// The resources stored in the `resources` slice are pointers to resource objects allocated by the engine.
// This is subtle and a little confusing. The reason for this is that the engine directly mutates resource objects
// that it creates and expects those mutations to be persisted directly to the snapshot.
type SnapshotManager struct {
	persister      SnapshotPersister    // The persister responsible for invalidating and persisting the snapshot
	baseSnapshot   *deploy.Snapshot     // The base snapshot for this plan
	secretsManager secrets.Manager      // The default secrets manager to use
	resources      []*resource.State    // The list of resources operated upon by this plan
	operations     []resource.Operation // The set of operations known to be outstanding in this plan

	// The set of resources that have been operated upon already by this plan. These resources could also have
	// been added to `resources` by other operations but need to be filtered out before writing the snapshot.
	dones map[*resource.State]bool

	completeOps      map[*resource.State]bool // The set of resources that have completed their operation
	mutationRequests chan<- mutationRequest   // The queue of mutation requests, to be retired serially by the manager
	cancel           chan bool                // A channel used to request cancellation of any new mutation requests.
	done             <-chan error             // A channel that sends a single result when the manager has shut down.

	refreshDeletes map[resource.URN]bool // The set of resources that have been deleted by a refresh in this plan.
}

var _ engine.SnapshotManager = (*SnapshotManager)(nil)

type mutationRequest struct {
	mutator func() bool
	result  chan<- error
}

func (sm *SnapshotManager) Close() error {
	close(sm.cancel)
	return <-sm.done
}

// If you need to understand what's going on in this file, start here!
//
// mutate is the serialization point for reads and writes of the global snapshot state.
// The given function will be, at the time of its invocation, the only function allowed to
// mutate state within the SnapshotManager.
//
// Serialization is performed by pushing the mutator function onto a channel, where another
// goroutine is polling the channel and executing the mutation functions as they come.
// This function optionally verifies the integrity of the snapshot before and after mutation.
//
// The mutator may indicate that its corresponding checkpoint write may be safely elided by
// returning `false`. As of this writing, we only elide writes after same steps with no
// meaningful changes (see sameSnapshotMutation.mustWrite for details). Any elided writes
// are flushed by the next non-elided write or the next call to Close.
//
// You should never observe or mutate the global snapshot without using this function unless
// you have a very good justification.
func (sm *SnapshotManager) mutate(mutator func() bool) error {
	result := make(chan error)
	select {
	case sm.mutationRequests <- mutationRequest{mutator: mutator, result: result}:
		return <-result
	case <-sm.cancel:
		return errors.New("snapshot manager closed")
	}
}

// RegisterResourceOutputs handles the registering of outputs on a Step that has already
// completed. This is accomplished by doing an in-place mutation of the resources currently
// resident in the snapshot.
//
// Due to the way this is currently implemented, the engine directly mutates output properties
// on the resource State object that it created. Since we are storing pointers to these objects
// in the `resources` slice, we need only to do a no-op mutation in order to flush these new
// mutations to disk.
//
// Note that this is completely not thread-safe and defeats the purpose of having a `mutate` callback
// entirely, but the hope is that this state of things will not be permament.
func (sm *SnapshotManager) RegisterResourceOutputs(step deploy.Step) error {
	return sm.mutate(func() bool {
		old, new := step.Old(), step.New()
		if old != nil && new != nil && old.Outputs.DeepEquals(new.Outputs) {
			logging.V(9).Infof("SnapshotManager: eliding RegisterResourceOutputs due to equal outputs")
			return false
		}

		return true
	})
}

// BeginMutation signals to the SnapshotManager that the engine intends to mutate the global snapshot
// by performing the given Step. This function gives the SnapshotManager a chance to record the
// intent to mutate before the mutation occurs.
func (sm *SnapshotManager) BeginMutation(step deploy.Step) (engine.SnapshotMutation, error) {
	contract.Requiref(step != nil, "step", "cannot be nil")
	logging.V(9).Infof("SnapshotManager: Beginning mutation for step `%s` on resource `%s`", step.Op(), step.URN())

	switch step.Op() {
	case deploy.OpSame:
		return &sameSnapshotMutation{sm}, nil
	case deploy.OpCreate, deploy.OpCreateReplacement:
		return sm.doCreate(step)
	case deploy.OpUpdate:
		return sm.doUpdate(step)
	case deploy.OpDelete, deploy.OpDeleteReplaced, deploy.OpReadDiscard, deploy.OpDiscardReplaced:
		return sm.doDelete(step)
	case deploy.OpReplace:
		return &replaceSnapshotMutation{sm}, nil
	case deploy.OpRead, deploy.OpReadReplacement:
		return sm.doRead(step)
	case deploy.OpRefresh:
		return &refreshSnapshotMutation{sm}, nil
	case deploy.OpRemovePendingReplace:
		return &removePendingReplaceSnapshotMutation{sm}, nil
	case deploy.OpImport, deploy.OpImportReplacement:
		return sm.doImport(step)
	}

	contract.Failf("unknown StepOp: %s", step.Op())
	return nil, nil
}

// All SnapshotMutation implementations in this file follow the same basic formula:
// mark the "old" state as done and mark the "new" state as new. The two special
// cases are Create (where the "old" state does not exist) and Delete (where the "new" state
// does not exist).
//
// Marking a resource state as old prevents it from being persisted to the snapshot in
// the `snap` function. Marking a resource state as new /enables/ it to be persisted to
// the snapshot in `snap`. See the comments in `snap` for more details.

type sameSnapshotMutation struct {
	manager *SnapshotManager
}

// mustWrite returns true if any semantically meaningful difference exists between the old and new states of a same
// step that forces us to write the checkpoint. If no such difference exists, the checkpoint write that corresponds to
// this step can be elided.
func (ssm *sameSnapshotMutation) mustWrite(step deploy.Step) bool {
	old := step.Old()
	new := step.New()

	contract.Assertf(old.Delete == new.Delete,
		"either both or neither resource must be pending deletion, got %v (old) != %v (new)",
		old.Delete, new.Delete)
	contract.Assertf(old.External == new.External,
		"either both or neither resource must be external, got %v (old) != %v (new)",
		old.External, new.External)

	if sameStep, isSameStep := step.(*deploy.SameStep); isSameStep {
		contract.Assertf(!sameStep.IsSkippedCreate(), "create cannot be skipped for SameStep")
	}

	// If the URN of this resource has changed, we must write the checkpoint. This should only be possible when a
	// resource is aliased.
	if old.URN != new.URN {
		logging.V(9).Infof("SnapshotManager: mustWrite() true because of URN")
		return true
	}

	// If the type of this resource has changed, we must write the checkpoint. This should only be possible when a
	// resource is aliased.
	if old.Type != new.Type {
		logging.V(9).Infof("SnapshotManager: mustWrite() true because of Type")
		return true
	}

	// If the kind of this resource has changed, we must write the checkpoint.
	if old.Custom != new.Custom {
		logging.V(9).Infof("SnapshotManager: mustWrite() true because of Custom")
		return true
	}

	// We need to persist the changes if CustomTimes have changed
	if old.CustomTimeouts != new.CustomTimeouts {
		logging.V(9).Infof("SnapshotManager: mustWrite() true because of CustomTimeouts")
		return true
	}

	// We need to persist the changes if CustomTimes have changed
	if old.RetainOnDelete != new.RetainOnDelete {
		logging.V(9).Infof("SnapshotManager: mustWrite() true because of RetainOnDelete")
		return true
	}

	contract.Assertf(old.ID == new.ID,
		"old and new resource IDs must be equal, got %v (old) != %v (new)", old.ID, new.ID)

	// If this resource's provider has changed, we must write the checkpoint. This can happen in scenarios involving
	// aliased providers or upgrades to default providers.
	if old.Provider != new.Provider {
		logging.V(9).Infof("SnapshotManager: mustWrite() true because of Provider")
		return true
	}

	// If this resource's parent has changed, we must write the checkpoint.
	if old.Parent != new.Parent {
		logging.V(9).Infof("SnapshotManager: mustWrite() true because of Parent")
		return true
	}

	// If the DeletedWith attribute of this resource has changed, we must write the checkpoint.
	if old.DeletedWith != new.DeletedWith {
		logging.V(9).Infof("SnapshotManager: mustWrite() true because of DeletedWith")
		return true
	}

	// If the protection attribute of this resource has changed, we must write the checkpoint.
	if old.Protect != new.Protect {
		logging.V(9).Infof("SnapshotManager: mustWrite() true because of Protect")
		return true
	}

	// If the inputs or outputs of this resource have changed, we must write the checkpoint. Note that it is possible
	// for the inputs of a "same" resource to have changed even if the contents of the input bags are different if the
	// resource's provider deems the physical change to be semantically irrelevant.
	if !old.Inputs.DeepEquals(new.Inputs) {
		logging.V(9).Infof("SnapshotManager: mustWrite() true because of Inputs")
		return true
	}
	if !old.Outputs.DeepEquals(new.Outputs) {
		logging.V(9).Infof("SnapshotManager: mustWrite() true because of Outputs")
		return true
	}

	// reflect.DeepEqual does not treat `nil` and `[]URN{}` as equal, so we must check for both
	// lists being empty ourselves.
	if len(old.Dependencies) != 0 || len(new.Dependencies) != 0 {
		// Sort dependencies before comparing them. If the dependencies have changed, we must write the checkpoint.
		sortDeps := func(deps []resource.URN) []resource.URN {
			result := make([]resource.URN, len(deps))
			copy(result, deps)
			slices.Sort(result)
			return result
		}
		oldDeps := sortDeps(old.Dependencies)
		newDeps := sortDeps(new.Dependencies)
		if !reflect.DeepEqual(oldDeps, newDeps) {
			logging.V(9).Infof("SnapshotManager: mustWrite() true because of Dependencies")
			return true
		}
	}

	if !reflect.DeepEqual(old.ResourceHooks, new.ResourceHooks) {
		logging.V(9).Infof("SnapshotManager: mustWrite() true because of ResourceHooks")
		return true
	}

	// Init errors are strictly advisory, so we do not consider them when deciding whether or not to write the
	// checkpoint. Likewise source positions are purely metadata and do not affect the system correctness, so
	// for performance we elide those as well. This prevents _every_ resource needing a snapshot write when
	// making large source code changes.

	logging.V(9).Infof("SnapshotManager: mustWrite() false")
	return false
}

func (ssm *sameSnapshotMutation) End(step deploy.Step, successful bool) error {
	contract.Requiref(step != nil, "step", "must not be nil")
	contract.Requiref(step.Op() == deploy.OpSame, "step.Op()", "must be %q, got %q", deploy.OpSame, step.Op())
	logging.V(9).Infof("SnapshotManager: sameSnapshotMutation.End(..., %v)", successful)
	return ssm.manager.mutate(func() bool {
		sameStep, isSameStep := step.(*deploy.SameStep)

		ssm.manager.markOperationComplete(step.New())
		if successful {
			ssm.manager.markDone(step.Old())

			// In the case of a 'resource create' in a program that wasn't specified by the user in the
			// --target list, we *never* want to write this to the checkpoint.  We treat it as if it
			// doesn't exist at all.  That way when the program runs the next time, we'll actually
			// create it.
			if isSameStep && sameStep.IsSkippedCreate() {
				return false
			}

			ssm.manager.markNew(step.New())

			// Note that "Same" steps only consider input and provider diffs, so it is possible to see a same step for a
			// resource with new dependencies, outputs, parent, protection. etc.
			//
			// As such, we diff all of the non-input properties of the resource here and write the snapshot if we find any
			// changes.
			if !ssm.mustWrite(step) {
				logging.V(9).Infof("SnapshotManager: sameSnapshotMutation.End() eliding write")
				return false
			}
		}

		logging.V(9).Infof("SnapshotManager: sameSnapshotMutation.End() not eliding write")
		return true
	})
}

func (sm *SnapshotManager) doCreate(step deploy.Step) (engine.SnapshotMutation, error) {
	logging.V(9).Infof("SnapshotManager.doCreate(%s)", step.URN())
	err := sm.mutate(func() bool {
		sm.markOperationPending(step.New(), resource.OperationTypeCreating)
		return true
	})
	if err != nil {
		return nil, err
	}

	return &createSnapshotMutation{sm}, nil
}

type createSnapshotMutation struct {
	manager *SnapshotManager
}

func (csm *createSnapshotMutation) End(step deploy.Step, successful bool) error {
	contract.Requiref(step != nil, "step", "must not be nil")
	logging.V(9).Infof("SnapshotManager: createSnapshotMutation.End(..., %v)", successful)
	return csm.manager.mutate(func() bool {
		csm.manager.markOperationComplete(step.New())
		if successful {
			// There is some very subtle behind-the-scenes magic here that
			// comes into play whenever this create is a CreateReplacement.
			//
			// Despite intending for the base snapshot to be immutable, the engine
			// does in fact mutate it by setting a `Delete` flag on resources
			// being replaced as part of a Create-Before-Delete replacement sequence.
			// Since we are storing the base snapshot and all resources by reference
			// (we have pointers to engine-allocated objects), this transparently
			// "just works" for the SnapshotManager.
			csm.manager.markNew(step.New())

			// If we had an old state that was marked as pending-replacement, mark its replacement as complete such
			// that it is flushed from the state file.
			if old := step.Old(); old != nil && old.PendingReplacement {
				csm.manager.markDone(old)
			}
		}
		return true
	})
}

func (sm *SnapshotManager) doUpdate(step deploy.Step) (engine.SnapshotMutation, error) {
	logging.V(9).Infof("SnapshotManager.doUpdate(%s)", step.URN())
	err := sm.mutate(func() bool {
		sm.markOperationPending(step.New(), resource.OperationTypeUpdating)
		return true
	})
	if err != nil {
		return nil, err
	}

	return &updateSnapshotMutation{sm}, nil
}

type updateSnapshotMutation struct {
	manager *SnapshotManager
}

func (usm *updateSnapshotMutation) End(step deploy.Step, successful bool) error {
	contract.Requiref(step != nil, "step", "must not be nil")
	logging.V(9).Infof("SnapshotManager: updateSnapshotMutation.End(..., %v)", successful)
	return usm.manager.mutate(func() bool {
		usm.manager.markOperationComplete(step.New())
		if successful {
			usm.manager.markDone(step.Old())
			usm.manager.markNew(step.New())
		}
		return true
	})
}

func (sm *SnapshotManager) doDelete(step deploy.Step) (engine.SnapshotMutation, error) {
	logging.V(9).Infof("SnapshotManager.doDelete(%s)", step.URN())
	err := sm.mutate(func() bool {
		sm.markOperationPending(step.Old(), resource.OperationTypeDeleting)
		return true
	})
	if err != nil {
		return nil, err
	}

	return &deleteSnapshotMutation{sm}, nil
}

type deleteSnapshotMutation struct {
	manager *SnapshotManager
}

func (dsm *deleteSnapshotMutation) End(step deploy.Step, successful bool) error {
	contract.Requiref(step != nil, "step", "must not be nil")
	logging.V(9).Infof("SnapshotManager: deleteSnapshotMutation.End(..., %v)", successful)
	return dsm.manager.mutate(func() bool {
		dsm.manager.markOperationComplete(step.Old())
		if successful {
			contract.Assertf(
				!step.Old().Protect ||
					step.Op() == deploy.OpDiscardReplaced ||
					step.Op() == deploy.OpDeleteReplaced,
				"Old must be unprotected (got %v) or the operation must be a replace (got %q)",
				step.Old().Protect, step.Op())

			if !step.Old().PendingReplacement {
				dsm.manager.markDone(step.Old())
			}
		}
		return true
	})
}

type replaceSnapshotMutation struct {
	manager *SnapshotManager
}

func (rsm *replaceSnapshotMutation) End(step deploy.Step, successful bool) error {
	logging.V(9).Infof("SnapshotManager: replaceSnapshotMutation.End(..., %v)", successful)
	return nil
}

func (sm *SnapshotManager) doRead(step deploy.Step) (engine.SnapshotMutation, error) {
	logging.V(9).Infof("SnapshotManager.doRead(%s)", step.URN())
	err := sm.mutate(func() bool {
		sm.markOperationPending(step.New(), resource.OperationTypeReading)
		return true
	})
	if err != nil {
		return nil, err
	}

	return &readSnapshotMutation{sm}, nil
}

type readSnapshotMutation struct {
	manager *SnapshotManager
}

func (rsm *readSnapshotMutation) End(step deploy.Step, successful bool) error {
	contract.Requiref(step != nil, "step", "must not be nil")
	logging.V(9).Infof("SnapshotManager: readSnapshotMutation.End(..., %v)", successful)
	return rsm.manager.mutate(func() bool {
		rsm.manager.markOperationComplete(step.New())
		if successful {
			if step.Old() != nil {
				rsm.manager.markDone(step.Old())
			}

			rsm.manager.markNew(step.New())
		}
		return true
	})
}

type refreshSnapshotMutation struct {
	manager *SnapshotManager
}

func (rsm *refreshSnapshotMutation) End(step deploy.Step, successful bool) error {
	contract.Requiref(step != nil, "step", "must not be nil")
	contract.Requiref(step.Op() == deploy.OpRefresh, "step.Op", "must be %q, got %q", deploy.OpRefresh, step.Op())
	logging.V(9).Infof("SnapshotManager: refreshSnapshotMutation.End(..., %v)", successful)
	return rsm.manager.mutate(func() bool {
		// We normally elide refreshes. The expectation is that all of these run before any actual mutations and that
		// some other component will rewrite the base snapshot in-memory, so there's no action the snapshot
		// manager needs to take other than to remember that the base snapshot--and therefore the actual snapshot--may
		// have changed.
		// The exception to this is persisted refreshes, which are not elided and are treated as normal operations.
		// These can either update or delete a resource.
		refreshStep, isRefreshStep := step.(*deploy.RefreshStep)
		viewStep, isViewStep := step.(*deploy.ViewStep)
		if (isViewStep && viewStep.Persisted()) || (isRefreshStep && refreshStep.Persisted()) {
			if successful {
				rsm.manager.markDone(step.Old())
				if step.New() != nil {
					rsm.manager.markNew(step.New())
				} else {
					rsm.manager.refreshDeletes[step.Old().URN] = true
				}
			}
			return true
		}

		return false
	})
}

type removePendingReplaceSnapshotMutation struct {
	manager *SnapshotManager
}

func (rsm *removePendingReplaceSnapshotMutation) End(step deploy.Step, successful bool) error {
	contract.Requiref(step != nil, "step", "must not be nil")
	contract.Requiref(step.Op() == deploy.OpRemovePendingReplace, "step.Op",
		"must be %q, got %q", deploy.OpRemovePendingReplace, step.Op())
	return rsm.manager.mutate(func() bool {
		res := step.Old()
		contract.Assertf(res.PendingReplacement, "resource %q must be pending replacement", res.URN)
		rsm.manager.markDone(res)
		return true
	})
}

func (sm *SnapshotManager) doImport(step deploy.Step) (engine.SnapshotMutation, error) {
	logging.V(9).Infof("SnapshotManager.doImport(%s)", step.URN())
	err := sm.mutate(func() bool {
		sm.markOperationPending(step.New(), resource.OperationTypeImporting)
		return true
	})
	if err != nil {
		return nil, err
	}

	return &importSnapshotMutation{sm}, nil
}

type importSnapshotMutation struct {
	manager *SnapshotManager
}

func (ism *importSnapshotMutation) End(step deploy.Step, successful bool) error {
	contract.Requiref(step != nil, "step", "must not be nil")
	contract.Requiref(step.Op() == deploy.OpImport || step.Op() == deploy.OpImportReplacement, "step.Op",
		"must be %q or %q, got %q", deploy.OpImport, deploy.OpImportReplacement, step.Op())

	return ism.manager.mutate(func() bool {
		ism.manager.markOperationComplete(step.New())
		if successful {
			ism.manager.markNew(step.New())
		}
		return true
	})
}

// markDone marks a resource as having been processed. Resources that have been marked
// in this manner won't be persisted in the snapshot.
func (sm *SnapshotManager) markDone(state *resource.State) {
	contract.Requiref(state != nil, "state", "must not be nil")
	sm.dones[state] = true
	logging.V(9).Infof("Marked old state snapshot as done: %v", state.URN)
}

// markNew marks a resource as existing in the new snapshot. This occurs on
// successful non-deletion operations where the given state is the new state
// of a resource that will be persisted to the snapshot.
func (sm *SnapshotManager) markNew(state *resource.State) {
	contract.Requiref(state != nil, "state", "must not be nil")
	sm.resources = append(sm.resources, state)
	logging.V(9).Infof("Appended new state snapshot to be written: %v", state.URN)
}

// markOperationPending marks a resource as undergoing an operation that will now be considered pending.
func (sm *SnapshotManager) markOperationPending(state *resource.State, op resource.OperationType) {
	contract.Requiref(state != nil, "state", "must not be nil")
	sm.operations = append(sm.operations, resource.NewOperation(state, op))
	logging.V(9).Infof("SnapshotManager.markPendingOperation(%s, %s)", state.URN, string(op))
}

// markOperationComplete marks a resource as having completed the operation that it previously was performing.
func (sm *SnapshotManager) markOperationComplete(state *resource.State) {
	contract.Requiref(state != nil, "state", "must not be nil")
	sm.completeOps[state] = true
	logging.V(9).Infof("SnapshotManager.markOperationComplete(%s)", state.URN)
}

// snap produces a new Snapshot given the base snapshot and a list of resources that the current
// plan has created.
func (sm *SnapshotManager) snap() *deploy.Snapshot {
	// At this point we have two resource DAGs. One of these is the base DAG for this plan; the other is the current DAG
	// for this plan. Any resource r may be present in both DAGs. In order to produce a snapshot, we need to merge these
	// DAGs such that all resource dependencies are correctly preserved. Conceptually, the merge proceeds as follows:
	//
	// - Begin with an empty merged DAG.
	// - For each resource r in the current DAG, insert r and its outgoing edges into the merged DAG.
	// - For each resource r in the base DAG:
	//     - If r is in the merged DAG, we are done: if the resource is in the merged DAG, it must have been in the
	//       current DAG, which accurately captures its current dependencies.
	//     - If r is not in the merged DAG, insert it and its outgoing edges into the merged DAG.
	//
	// Physically, however, each DAG is represented as list of resources without explicit dependency edges. In place of
	// edges, it is assumed that the list represents a valid topological sort of its source DAG. Thus, any resource r at
	// index i in a list L must be assumed to be dependent on all resources in L with index j s.t. j < i. Due to this
	// representation, we implement the algorithm above as follows to produce a merged list that represents a valid
	// topological sort of the merged DAG:
	//
	// - Begin with an empty merged list.
	// - For each resource r in the current list, append r to the merged list. r must be in a correct location in the
	//   merged list, as its position relative to its assumed dependencies has not changed.
	// - For each resource r in the base list:
	//     - If r is in the merged list, we are done by the logic given in the original algorithm.
	//     - If r is not in the merged list, append r to the merged list. r must be in a correct location in the merged
	//       list:
	//         - If any of r's dependencies were in the current list, they must already be in the merged list and their
	//           relative order w.r.t. r has not changed.
	//         - If any of r's dependencies were not in the current list, they must already be in the merged list, as
	//           they would have been appended to the list before r.

	// Start with a copy of the resources produced during the evaluation of the current plan.
	resources := make([]*resource.State, 0, len(sm.resources))

	// If any resources are "done", we need to filter them out here. These could be resources that have been later
	// deleted, or had some other operation performed on them such as an import then an update.
	for _, res := range sm.resources {
		if !sm.dones[res] {
			resources = append(resources, res)
		}
	}

	// Append any resources from the base plan that were not produced by the current plan.
	if base := sm.baseSnapshot; base != nil {
		for _, res := range base.Resources {
			if !sm.dones[res] {
				resources = append(resources, res)
			}
		}
	}

	// Filter any refresh deletes
	engine.FilterRefreshDeletes(sm.refreshDeletes, resources)

	// Record any pending operations, if there are any outstanding that have not completed yet.
	var operations []resource.Operation
	for _, op := range sm.operations {
		if !sm.completeOps[op.Resource] {
			operations = append(operations, op)
		}
	}

	// Track pending create operations from the base snapshot
	// and propagate them to the new snapshot: we don't want to clear pending CREATE operations
	// because these must require user intervention to be cleared or resolved.
	if base := sm.baseSnapshot; base != nil {
		for _, pendingOperation := range base.PendingOperations {
			if pendingOperation.Type == resource.OperationTypeCreating {
				operations = append(operations, pendingOperation)
			}
		}
	}

	manifest := deploy.Manifest{
		Time:    time.Now(),
		Version: version.Version,
		// Plugins: sm.plugins, - Explicitly dropped, since we don't use the plugin list in the manifest anymore.
	}

	// The backend.SnapshotManager and backend.SnapshotPersister will keep track of any changes to
	// the Snapshot (checkpoint file) in the HTTP backend. We will reuse the snapshot's secrets manager when possible
	// to ensure that secrets are not re-encrypted on each update.
	secretsManager := sm.secretsManager
	if sm.baseSnapshot != nil && secrets.AreCompatible(secretsManager, sm.baseSnapshot.SecretsManager) {
		secretsManager = sm.baseSnapshot.SecretsManager
	}

	var metadata deploy.SnapshotMetadata
	if sm.baseSnapshot != nil {
		metadata = sm.baseSnapshot.Metadata
	}

	manifest.Magic = manifest.NewMagic()
	return deploy.NewSnapshot(manifest, secretsManager, resources, operations, metadata)
}

// saveSnapshot persists the current snapshot. If integrity checking is enabled,
// the snapshot's integrity is also verified. If the snapshot is invalid,
// metadata about this write operation is added to the snapshot before it is
// written, in order to aid debugging should future operations fail with an
// error.
func (sm *SnapshotManager) saveSnapshot() error {
	snap, err := sm.snap().NormalizeURNReferences()
	if err != nil {
		return fmt.Errorf("failed to normalize URN references: %w", err)
	}

	// In order to persist metadata about snapshot integrity issues, we check the
	// snapshot's validity *before* we write it. However, should an error occur,
	// we will only raise this *after* the write has completed. In the event that
	// integrity checking is disabled, we still actually perform the check (and
	// write metadata appropriately), but we will not raise the error following a
	// successful write.
	//
	// If the actual write fails for any reason, this error will supersede any
	// integrity error. This matches behaviour prior to when integrity metadata
	// writing was introduced.
	//
	// Metadata will be cleared out by a successful operation (even if integrity
	// checking is being enforced).
	integrityError := snap.VerifyIntegrity()
	if integrityError == nil {
		snap.Metadata.IntegrityErrorMetadata = nil
	} else {
		snap.Metadata.IntegrityErrorMetadata = &deploy.SnapshotIntegrityErrorMetadata{
			Version: version.Version,
			Command: strings.Join(os.Args, " "),
			Error:   integrityError.Error(),
		}
	}

	if err := sm.persister.Save(snap); err != nil {
		return fmt.Errorf("failed to save snapshot: %w", err)
	}
	if !DisableIntegrityChecking && integrityError != nil {
		return fmt.Errorf("failed to verify snapshot: %w", integrityError)
	}
	return nil
}

// defaultServiceLoop saves a Snapshot whenever a mutation occurs
func (sm *SnapshotManager) defaultServiceLoop(mutationRequests chan mutationRequest, done chan error) {
	// True if we have elided writes since the last actual write.
	hasElidedWrites := true

	// Service each mutation request in turn.
serviceLoop:
	for {
		select {
		case request := <-mutationRequests:
			var err error
			if request.mutator() {
				err = sm.saveSnapshot()
				hasElidedWrites = false
			} else {
				hasElidedWrites = true
			}
			request.result <- err
		case <-sm.cancel:
			break serviceLoop
		}
	}

	// If we still have elided writes once the channel has closed, flush the snapshot.
	var err error
	if hasElidedWrites {
		logging.V(9).Infof("SnapshotManager: flushing elided writes...")
		err = sm.saveSnapshot()
	}
	done <- err
}

// unsafeServiceLoop doesn't save Snapshots when mutations occur and instead saves Snapshots when
// SnapshotManager.Close() is invoked. It trades reliability for speed as every mutation does not
// cause a Snapshot to be serialized to the user's state backend.
func (sm *SnapshotManager) unsafeServiceLoop(mutationRequests chan mutationRequest, done chan error) {
	for {
		select {
		case request := <-mutationRequests:
			request.mutator()
			request.result <- nil
		case <-sm.cancel:
			done <- sm.saveSnapshot()
			return
		}
	}
}

// NewSnapshotManager creates a new SnapshotManager for the given stack name, using the given persister, default secrets
// manager and base snapshot.
//
// It is *very important* that the baseSnap pointer refers to the same Snapshot given to the engine! The engine will
// mutate this object and correctness of the SnapshotManager depends on being able to observe this mutation. (This is
// not ideal...)
func NewSnapshotManager(
	persister SnapshotPersister,
	secretsManager secrets.Manager,
	baseSnap *deploy.Snapshot,
) *SnapshotManager {
	mutationRequests, cancel, done := make(chan mutationRequest), make(chan bool), make(chan error)

	manager := &SnapshotManager{
		persister:        persister,
		secretsManager:   secretsManager,
		baseSnapshot:     baseSnap,
		dones:            make(map[*resource.State]bool),
		completeOps:      make(map[*resource.State]bool),
		mutationRequests: mutationRequests,
		cancel:           cancel,
		done:             done,
		refreshDeletes:   make(map[resource.URN]bool),
	}

	serviceLoop := manager.defaultServiceLoop

	if env.SkipCheckpoints.Value() {
		serviceLoop = manager.unsafeServiceLoop
	}

	go serviceLoop(mutationRequests, done)

	return manager
}
