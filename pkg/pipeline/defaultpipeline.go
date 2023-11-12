package pipeline

import (
	"fmt"
	"github.com/cloudintegrator/go-sonic-core/api/core"
	"log/slog"
)

var PIPELINE []interface{}

type DefaultPipeline struct {
	SOPath   string
	SonicApp *core.SonicApp
}

// ParseSonicApp reads yaml application i.e SonicApp.
func (p *DefaultPipeline) ParseSonicApp() {
	slog.Debug("********** App starting **********" + p.SonicApp.App.Name)
	for i := 0; i < len(p.SonicApp.App.Flows); i++ {
		f := p.SonicApp.App.Flows[i]
		p.InitializeFlowComponents(&f)
	}
}

// InitializeFlowComponents reads Flow element of SonicApp and traverses its Component.
func (p *DefaultPipeline) InitializeFlowComponents(flow *core.Flow) {
	if len(flow.Components) != 0 {
		for i := 0; i < len(flow.Components); i++ {
			// Get the component from flow.
			component := flow.Components[i]
			slog.Debug("********** Initializing Component **********: " + component.Name)

			// Get the component configuration from Component.
			componentConfig := component.Configuration
			mapConfig := componentConfig.(map[string]interface{})
			value := mapConfig["ref"]
			if value != nil {
				componentConfig = p.GetReferenceConfiguration(value.(string))
			}

			// Build the pipeline with Component & Component Configuration.
			p.BuildPipeline(component, componentConfig)
			slog.Debug("*********** " + value.(string))
		}
	}
}

// GetReferenceConfiguration reads the "ref" element (if present) in the Configuration of the Component
func (p *DefaultPipeline) GetReferenceConfiguration(component string) (out interface{}) {
	slog.Debug("********** Get reference configuration for:" + component)
	if len(p.SonicApp.App.Common) != 0 {
		for i := 0; i < len(p.SonicApp.App.Common); i++ {
			config := p.SonicApp.App.Common[0]
			if component == config.Name {
				return config
			}
		}
	}
	return nil
}

// BuildPipeline builds the actual pipeline with the Component extracted from Flow.
// Based on its Kind property it will find the appropriate so (shared object) file and instantiate the components i.e
// SonicSource, SonicProcessor etc.
func (p *DefaultPipeline) BuildPipeline(component core.Component, component_config interface{}) {
	fmt.Println("TODO")
}
