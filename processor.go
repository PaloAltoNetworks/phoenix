// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package phoenix

import (
	"encoding/json"

	"go.aporeto.io/bahamut"
	"go.aporeto.io/gaia"
)

type remoteProcessorProcessor struct {
	pluginsRegistry HooksRegistry
}

func newRemoteProcessorProcessor(pluginsRegistry HooksRegistry) *remoteProcessorProcessor {

	return &remoteProcessorProcessor{
		pluginsRegistry: pluginsRegistry,
	}
}

// ProcessCreate is triggered with a new plugin
func (p *remoteProcessorProcessor) ProcessCreate(ctx bahamut.Context) error {

	// Retrieve input data
	rp := ctx.InputData().(*gaia.RemoteProcessor)
	obj := gaia.Manager().IdentifiableFromString(rp.TargetIdentity)

	if err := json.Unmarshal(rp.Input, &obj); err != nil {
		return err
	}

	for _, pluginFunc := range p.pluginsRegistry {

		if err := pluginFunc(rp.RequestID, rp.Operation, rp.Mode, obj, rp.Claims); err != nil {
			return err
		}
	}

	rp.Output = obj

	ctx.SetOutputData(rp)

	return nil
}
