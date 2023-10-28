package pipeline

import (
	"github.com/cloudintegrator/go-sonic-core/api/core"
	"log/slog"
)

func BuildPipeline(a *core.SonicApp) {
	slog.Debug("********** App starting ********** %s", a.App.Name, nil)
	for i := 0; i < len(a.App.Flows); i++ {
		f := a.App.Flows[i]
		InitializeFlowComponents(&f)
	}
}

func InitializeFlowComponents(flow *core.Flow) {
	if len(flow.Components) != 0 {
		for i := 0; i < len(flow.Components); i++ {
			component := flow.Components[i]
			slog.Debug("********** Initilizing Component ********** %s", component.Name, nil)

			component_config := component.Configuration

			slog.Debug("********** Initilizing Component Configuratio ********** %s", component_config, nil)
		}
	}
}

func GetReferenceConfiguration(component string, a *core.SonicApp) (out interface{}) {
	slog.Debug("********** Get reference configuration for: %s", component, nil)
	if len(a.App.Common) != 0 {
		for i := 0; i < len(a.App.Common); i++ {
			config := a.App.Common[0]
			if component == config.Name {
				return config
			}
		}
	}
	return nil
}
