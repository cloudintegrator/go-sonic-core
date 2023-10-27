package pipeline

import (
	"github.com/cloudintegrator/go-sonic-core/api/core"
	"log"
)

func BuildPipeline(a *core.SonicApp) {
	log.Printf("********** App starting ********** %s", a.App.Name)
	for i := 0; i < len(a.App.Flows); i++ {
		f := a.App.Flows[i]
		InitializeFlowComponents(&f)
	}
}

func InitializeFlowComponents(flow *core.Flow) {
	if len(flow.Components) != 0 {
		for i := 0; i < len(flow.Components); i++ {
			component := flow.Components[i]
			log.Printf("********** Initilizing Component ********** %s", component.Name)

			component_config := component.Configuration
			log.Printf("********** Initilizing Component Configuratio ********** %s", component_config)
		}
	}
}
