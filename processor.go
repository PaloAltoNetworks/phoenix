package phoenix

import (
	"encoding/json"

	"github.com/aporeto-inc/bahamut"
	"github.com/aporeto-inc/gaia/rufusmodels/v1/golang"
	"github.com/aporeto-inc/gaia/squallmodels/v1/golang"
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
func (p *remoteProcessorProcessor) ProcessCreate(ctx *bahamut.Context) error {

	// Retrieve input data
	rp := ctx.InputData.(*rufusmodels.RemoteProcessor)
	obj := squallmodels.IdentifiableForIdentity(rp.TargetIdentity)

	if err := json.Unmarshal(rp.Input, &obj); err != nil {
		return err
	}

	for _, pluginFunc := range p.pluginsRegistry {

		if err := pluginFunc(rp.RequestID, rp.Operation, rp.Mode, obj, rp.Claims); err != nil {
			return err
		}
	}

	rp.Output = obj

	ctx.OutputData = rp

	return nil
}
