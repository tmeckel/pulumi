// Copyright 2016-2018, Pulumi Corporation.
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

package deploytest

import (
	"context"
	"errors"

	"github.com/blang/semver"
	uuid "github.com/gofrs/uuid"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/pulumi/sdk/v3/go/common/workspace"
)

type Provider struct {
	plugin.NotForwardCompatibleProvider

	Name    string
	Package tokens.Package
	Version semver.Version

	Config     resource.PropertyMap
	configured bool

	DialMonitorF func(ctx context.Context, endpoint string) (*ResourceMonitor, error)

	ParameterizeF func(ctx context.Context, request plugin.ParameterizeRequest) (plugin.ParameterizeResponse, error)

	GetSchemaF func(request plugin.GetSchemaRequest) ([]byte, error)

	CheckConfigF func(urn resource.URN, olds,
		news resource.PropertyMap, allowUnknowns bool) (resource.PropertyMap, []plugin.CheckFailure, error)
	DiffConfigF func(urn resource.URN, oldInputs, oldOutputs, newInputs resource.PropertyMap,
		ignoreChanges []string) (plugin.DiffResult, error)
	ConfigureF func(news resource.PropertyMap) error

	CheckF func(urn resource.URN,
		olds, news resource.PropertyMap, randomSeed []byte) (resource.PropertyMap, []plugin.CheckFailure, error)
	DiffF func(urn resource.URN, id resource.ID, oldInputs, oldOutputs, newInputs resource.PropertyMap,
		ignoreChanges []string) (plugin.DiffResult, error)
	CreateF func(urn resource.URN, inputs resource.PropertyMap, timeout float64,
		preview bool) (resource.ID, resource.PropertyMap, resource.Status, error)
	UpdateF func(urn resource.URN, id resource.ID, oldInputs, oldOutputs, newInputs resource.PropertyMap, timeout float64,
		ignoreChanges []string, preview bool) (resource.PropertyMap, resource.Status, error)
	DeleteF func(urn resource.URN, id resource.ID,
		oldInputs, oldOutputs resource.PropertyMap, timeout float64) (resource.Status, error)
	ReadF func(urn resource.URN, id resource.ID,
		inputs, state resource.PropertyMap) (plugin.ReadResult, resource.Status, error)

	ConstructF func(monitor *ResourceMonitor, typ, name string, parent resource.URN, inputs resource.PropertyMap,
		info plugin.ConstructInfo, options plugin.ConstructOptions) (plugin.ConstructResult, error)

	InvokeF func(tok tokens.ModuleMember,
		inputs resource.PropertyMap) (resource.PropertyMap, []plugin.CheckFailure, error)
	StreamInvokeF func(tok tokens.ModuleMember, args resource.PropertyMap,
		onNext func(resource.PropertyMap) error) ([]plugin.CheckFailure, error)

	CallF func(monitor *ResourceMonitor, tok tokens.ModuleMember, args resource.PropertyMap, info plugin.CallInfo,
		options plugin.CallOptions) (plugin.CallResult, error)

	CancelF func() error

	GetMappingF  func(key, provider string) ([]byte, string, error)
	GetMappingsF func(key string) ([]string, error)
}

func (prov *Provider) SignalCancellation() error {
	if prov.CancelF == nil {
		return nil
	}
	return prov.CancelF()
}

func (prov *Provider) Close() error {
	return nil
}

func (prov *Provider) Pkg() tokens.Package {
	return prov.Package
}

func (prov *Provider) GetPluginInfo() (workspace.PluginInfo, error) {
	return workspace.PluginInfo{
		Name:    prov.Name,
		Version: &prov.Version,
	}, nil
}

func (prov *Provider) Parameterize(
	ctx context.Context, params plugin.ParameterizeRequest,
) (plugin.ParameterizeResponse, error) {
	if prov.ParameterizeF == nil {
		return plugin.ParameterizeResponse{}, errors.New("no parameters")
	}
	return prov.ParameterizeF(ctx, params)
}

func (prov *Provider) GetSchema(request plugin.GetSchemaRequest) ([]byte, error) {
	if prov.GetSchemaF == nil {
		return []byte("{}"), nil
	}
	return prov.GetSchemaF(request)
}

func (prov *Provider) CheckConfig(urn resource.URN, olds,
	news resource.PropertyMap, allowUnknowns bool,
) (resource.PropertyMap, []plugin.CheckFailure, error) {
	if prov.CheckConfigF == nil {
		return news, nil, nil
	}
	return prov.CheckConfigF(urn, olds, news, allowUnknowns)
}

func (prov *Provider) DiffConfig(urn resource.URN, oldInputs, oldOutputs, newInputs resource.PropertyMap, _ bool,
	ignoreChanges []string,
) (plugin.DiffResult, error) {
	if prov.DiffConfigF == nil {
		return plugin.DiffResult{}, nil
	}
	return prov.DiffConfigF(urn, oldInputs, oldOutputs, newInputs, ignoreChanges)
}

func (prov *Provider) Configure(inputs resource.PropertyMap) error {
	contract.Assertf(!prov.configured, "provider %v was already configured", prov.Name)
	prov.configured = true

	if prov.ConfigureF == nil {
		prov.Config = inputs
		return nil
	}
	return prov.ConfigureF(inputs)
}

func (prov *Provider) Check(urn resource.URN,
	olds, news resource.PropertyMap, _ bool, randomSeed []byte,
) (resource.PropertyMap, []plugin.CheckFailure, error) {
	contract.Requiref(randomSeed != nil, "randomSeed", "must not be nil")
	if prov.CheckF == nil {
		return news, nil, nil
	}
	return prov.CheckF(urn, olds, news, randomSeed)
}

func (prov *Provider) Create(urn resource.URN, props resource.PropertyMap, timeout float64,
	preview bool,
) (resource.ID, resource.PropertyMap, resource.Status, error) {
	if prov.CreateF == nil {
		// generate a new uuid
		uuid, err := uuid.NewV4()
		if err != nil {
			return "", nil, resource.StatusOK, err
		}
		return resource.ID(uuid.String()), resource.PropertyMap{}, resource.StatusOK, nil
	}
	return prov.CreateF(urn, props, timeout, preview)
}

func (prov *Provider) Diff(urn resource.URN, id resource.ID,
	oldInputs, oldOutputs, newInputs resource.PropertyMap, _ bool, ignoreChanges []string,
) (plugin.DiffResult, error) {
	if prov.DiffF == nil {
		return plugin.DiffResult{}, nil
	}
	return prov.DiffF(urn, id, oldInputs, oldOutputs, newInputs, ignoreChanges)
}

func (prov *Provider) Update(urn resource.URN, id resource.ID, oldInputs, oldOutputs, newInputs resource.PropertyMap,
	timeout float64, ignoreChanges []string, preview bool,
) (resource.PropertyMap, resource.Status, error) {
	if prov.UpdateF == nil {
		return newInputs, resource.StatusOK, nil
	}
	return prov.UpdateF(urn, id, oldInputs, oldOutputs, newInputs, timeout, ignoreChanges, preview)
}

func (prov *Provider) Delete(urn resource.URN,
	id resource.ID, oldInputs, oldOutputs resource.PropertyMap, timeout float64,
) (resource.Status, error) {
	if prov.DeleteF == nil {
		return resource.StatusOK, nil
	}
	return prov.DeleteF(urn, id, oldInputs, oldOutputs, timeout)
}

func (prov *Provider) Read(urn resource.URN, id resource.ID,
	inputs, state resource.PropertyMap,
) (plugin.ReadResult, resource.Status, error) {
	contract.Assertf(urn != "", "Read URN was empty")
	contract.Assertf(id != "", "Read ID was empty")
	if prov.ReadF == nil {
		return plugin.ReadResult{
			Outputs: resource.PropertyMap{},
			Inputs:  resource.PropertyMap{},
		}, resource.StatusUnknown, nil
	}
	return prov.ReadF(urn, id, inputs, state)
}

func (prov *Provider) Construct(info plugin.ConstructInfo, typ tokens.Type, name string, parent resource.URN,
	inputs resource.PropertyMap, options plugin.ConstructOptions,
) (plugin.ConstructResult, error) {
	if prov.ConstructF == nil {
		return plugin.ConstructResult{}, nil
	}
	dialMonitorImpl := dialMonitor
	if prov.DialMonitorF != nil {
		dialMonitorImpl = prov.DialMonitorF
	}
	monitor, err := dialMonitorImpl(context.Background(), info.MonitorAddress)
	if err != nil {
		return plugin.ConstructResult{}, err
	}
	return prov.ConstructF(monitor, string(typ), name, parent, inputs, info, options)
}

func (prov *Provider) Invoke(tok tokens.ModuleMember,
	args resource.PropertyMap,
) (resource.PropertyMap, []plugin.CheckFailure, error) {
	if prov.InvokeF == nil {
		return resource.PropertyMap{}, nil, nil
	}
	return prov.InvokeF(tok, args)
}

func (prov *Provider) StreamInvoke(
	tok tokens.ModuleMember, args resource.PropertyMap,
	onNext func(resource.PropertyMap) error,
) ([]plugin.CheckFailure, error) {
	if prov.StreamInvokeF == nil {
		return []plugin.CheckFailure{}, errors.New("StreamInvoke unimplemented")
	}
	return prov.StreamInvokeF(tok, args, onNext)
}

func (prov *Provider) Call(tok tokens.ModuleMember, args resource.PropertyMap, info plugin.CallInfo,
	options plugin.CallOptions,
) (plugin.CallResult, error) {
	if prov.CallF == nil {
		return plugin.CallResult{}, nil
	}
	dialMonitorImpl := dialMonitor
	if prov.DialMonitorF != nil {
		dialMonitorImpl = prov.DialMonitorF
	}
	monitor, err := dialMonitorImpl(context.Background(), info.MonitorAddress)
	if err != nil {
		return plugin.CallResult{}, err
	}
	return prov.CallF(monitor, tok, args, info, options)
}

func (prov *Provider) GetMapping(key, provider string) ([]byte, string, error) {
	if prov.GetMappingF == nil {
		return nil, "", nil
	}
	return prov.GetMappingF(key, provider)
}

func (prov *Provider) GetMappings(key string) ([]string, error) {
	if prov.GetMappingsF == nil {
		return []string{}, nil
	}
	return prov.GetMappingsF(key)
}
