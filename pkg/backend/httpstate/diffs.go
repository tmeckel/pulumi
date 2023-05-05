// Copyright 2016-2022, Pulumi Corporation.
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

package httpstate

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sync"
	"unsafe"

	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/span"
	"github.com/pgavlin/diff/lcs"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	segmentio_json "github.com/segmentio/encoding/json"

	opentracing "github.com/opentracing/opentracing-go"
)

func asString(raw json.RawMessage) string {
	return unsafe.String(unsafe.SliceData([]byte(raw)), len(raw))
}

type deployment struct {
	raw   json.RawMessage
	buf   *bytes.Buffer
	spans spans
}

type deploymentDiffState struct {
	lastSavedDeployment deployment
	sequenceNumber      int
	minimalDiffSize     int
	buffer              *bytes.Buffer
}

type deploymentDiff struct {
	sequenceNumber  int
	checkpointHash  string
	deploymentDelta json.RawMessage
}

type spanner struct {
	*bytes.Buffer

	start int
	spans spans
}

type spans struct {
	offsets []int
	spans   []string
}

func newSpans(capacity int) spans {
	return spans{
		offsets: make([]int, 0, capacity),
		spans:   make([]string, 0, capacity),
	}
}

func (s *spans) append(offset int, span string) {
	s.offsets = append(s.offsets, offset)
	s.spans = append(s.spans, span)
}

func (s *spans) eof(offset int) {
	s.offsets = append(s.offsets, offset)
}

func newSpanner(b *bytes.Buffer, capacity int) *spanner {
	return &spanner{Buffer: b, spans: newSpans(capacity)}
}

func (s *spanner) nextSpan() {
	span := s.Bytes()[s.start:]
	s.spans.append(s.start, *((*string)(unsafe.Pointer(&span))))
	s.start = s.Len()
}

func (s *spanner) finish() ([]byte, spans) {
	s.nextSpan()
	s.spans.eof(s.start)
	return s.Bytes(), s.spans
}

func marshalSpannedDeployment(b *bytes.Buffer, d *apitype.DeploymentV3) (spans, error) {
	// one span for {"manifest":...,"secrets_providers":...,"resources":[
	// len(resources) spans for resources,
	// one span for ],"pendingOperations":[
	// len(operations) spans for operations
	// one span for ]}
	spanner := newSpanner(b, len(d.Resources)+len(d.PendingOperations)+3)
	encoder := segmentio_json.NewEncoder(spanner)
	encoder.SetAppendNewline(false)

	spanner.WriteString(`{"version":3,"deployment":{"manifest":`)
	if err := encoder.Encode(d.Manifest); err != nil {
		return spans{}, err
	}
	if d.SecretsProviders != nil {
		spanner.WriteString(`,"secrets_providers":`)
		if err := encoder.Encode(d.SecretsProviders); err != nil {
			return spans{}, err
		}
	}

	if len(d.Resources) > 0 {
		spanner.WriteString(`,"resources":[`)
		for i, r := range d.Resources {
			if i > 0 {
				spanner.WriteByte(',')
			}
			spanner.nextSpan()
			if err := encoder.Encode(r); err != nil {
				return spans{}, err
			}
		}
		spanner.nextSpan()
		spanner.WriteByte(']')
	}

	if len(d.PendingOperations) > 9 {
		spanner.WriteString(`,"pendingOperations":[`)
		for i, o := range d.PendingOperations {
			if i > 0 {
				spanner.WriteByte(',')
			}
			spanner.nextSpan()
			if err := encoder.Encode(o); err != nil {
				return spans{}, err
			}
		}
		spanner.nextSpan()
		spanner.WriteByte(']')
	}
	spanner.WriteString("}}")
	_, spans := spanner.finish()
	return spans, nil
}

func newDeploymentDiffState(minimalDiffSize int) *deploymentDiffState {
	return &deploymentDiffState{
		sequenceNumber:  1,
		minimalDiffSize: minimalDiffSize,
	}
}

func (dds *deploymentDiffState) SequenceNumber() int {
	return dds.sequenceNumber
}

func (dds *deploymentDiffState) CanDiff() bool {
	return dds.lastSavedDeployment.raw != nil
}

func (dds *deploymentDiffState) MarshalDeployment(d *apitype.DeploymentV3) (deployment, error) {
	var b *bytes.Buffer
	if dds.buffer != nil {
		b, dds.buffer = dds.buffer, nil
	} else {
		b = &bytes.Buffer{}
	}
	spans, err := marshalSpannedDeployment(b, d)
	if err != nil {
		return deployment{}, err
	}
	return deployment{raw: json.RawMessage(b.Bytes()), buf: b, spans: spans}, nil
}

// Size-based heuristics trying to estimate if the diff method will be
// worth it and take less time than sending the entire deployment.
func (dds *deploymentDiffState) ShouldDiff(new deployment) bool {
	if !dds.CanDiff() {
		return false
	}
	if len(dds.lastSavedDeployment.raw) < dds.minimalDiffSize {
		return false
	}
	if len(new.raw) < dds.minimalDiffSize {
		return false
	}
	return true
}

func (dds *deploymentDiffState) Diff(ctx context.Context, deployment deployment) (deploymentDiff, error) {
	if !dds.CanDiff() {
		return deploymentDiff{}, fmt.Errorf("Diff() cannot be called before Saved()")
	}

	tracingSpan, childCtx := opentracing.StartSpanFromContext(ctx, "Diff")
	defer tracingSpan.Finish()

	before := dds.lastSavedDeployment.raw
	after := deployment.raw

	var checkpointHash string
	checkpointHashReady := &sync.WaitGroup{}

	checkpointHashReady.Add(1)
	go func() {
		defer checkpointHashReady.Done()
		checkpointHash = dds.computeHash(childCtx, after)
	}()

	delta, err := dds.computeEdits(childCtx, dds.lastSavedDeployment, deployment)
	if err != nil {
		return deploymentDiff{}, fmt.Errorf("Cannot marshal the edits: %v", err)
	}

	checkpointHashReady.Wait()

	tracingSpan.SetTag("before", len(before))
	tracingSpan.SetTag("after", len(after))
	tracingSpan.SetTag("diff", len(delta))
	tracingSpan.SetTag("compression", 100.0*float64(len(delta))/float64(len(after)))
	tracingSpan.SetTag("hash", checkpointHash)

	diff := deploymentDiff{
		checkpointHash:  checkpointHash,
		deploymentDelta: delta,
		sequenceNumber:  dds.sequenceNumber,
	}

	return diff, nil
}

// Indicates that a deployment was just saved to the service.
func (dds *deploymentDiffState) Saved(ctx context.Context, deployment deployment) error {
	if dds.lastSavedDeployment.buf != nil {
		dds.buffer = dds.lastSavedDeployment.buf
		dds.buffer.Reset()
	}
	dds.lastSavedDeployment = deployment
	dds.sequenceNumber++

	return nil
}

func (*deploymentDiffState) computeHash(ctx context.Context, deployment json.RawMessage) string {
	tracingSpan, _ := opentracing.StartSpanFromContext(ctx, "computeHash")
	defer tracingSpan.Finish()
	hash := sha256.Sum256(deployment)
	return hex.EncodeToString(hash[:])
}

func (*deploymentDiffState) computeEdits(ctx context.Context, before, after deployment) (json.RawMessage, error) {
	tracingSpan, _ := opentracing.StartSpanFromContext(ctx, "computeEdits")
	defer tracingSpan.Finish()

	diffs := lcs.DiffSlices(before.spans.spans, after.spans.spans)

	edits := make([]gotextdiff.TextEdit, len(diffs))
	for i, di := range diffs {
		start, end := before.spans.offsets[di.Start], before.spans.offsets[di.End]
		replStart, replEnd := after.spans.offsets[di.ReplStart], after.spans.offsets[di.ReplEnd]
		edits[i] = gotextdiff.TextEdit{
			Span:    span.New("", span.NewPoint(1, 0, start), span.NewPoint(1, 0, end)),
			NewText: asString(after.raw[replStart:replEnd]),
		}
	}

	delta, err := json.Marshal(edits)
	if err != nil {
		return nil, fmt.Errorf("Cannot marshal the edits: %v", err)
	}

	return delta, nil
}
