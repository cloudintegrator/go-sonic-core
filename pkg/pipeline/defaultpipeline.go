package pipeline

import (
	"github.com/cloudintegrator/go-sonic-core/api/core"
	"log/slog"
)

var pipeline []interface{}

func ParseSonicApp(a *core.SonicApp) {
	slog.Debug("********** App starting **********" + a.App.Name)
	for i := 0; i < len(a.App.Flows); i++ {
		f := a.App.Flows[i]
		InitializeFlowComponents(a, &f)
	}
}

func InitializeFlowComponents(a *core.SonicApp, flow *core.Flow) {
	if len(flow.Components) != 0 {
		for i := 0; i < len(flow.Components); i++ {
			// Get the component from
			component := flow.Components[i]
			slog.Debug("********** Initilizing Component **********: " + component.Name)

			// Get the component configuration
			component_config := component.Configuration
			mapConfig := component_config.(map[string]interface{})
			value := mapConfig["ref"]
			if value != nil {
				component_config = GetReferenceConfiguration(value.(string), a)
			}

			slog.Debug("*********** " + value.(string))
		}
	}
}

func GetReferenceConfiguration(component string, a *core.SonicApp) (out interface{}) {
	slog.Debug("********** Get reference configuration for:" + component)
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
