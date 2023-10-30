package pipeline

import (
	"github.com/cloudintegrator/go-sonic-core/api/core"
	"github.com/cloudintegrator/go-sonic-core/pkg/utility"
	"testing"
)

func TestParseSonicApp(t *testing.T) {
	c := &core.SonicApp{}
	utility.YamlToStruct("../../spec/app-spec.yaml", c)

	p := DefaultPipeline{
		SOPath:   "/Users/anupam.gogoi.br/github/anupamgogoi/sonic/go-sonic-kafka-connector/plugin/",
		SonicApp: c,
	}
	p.ParseSonicApp()
}
