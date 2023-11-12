package pipeline

import (
	"github.com/cloudintegrator/go-sonic-core/api/core"
	"github.com/cloudintegrator/go-sonic-core/pkg/utility"
	"testing"
)

func TestParseSonicAppFromFile(t *testing.T) {
	c := &core.SonicApp{}
	err := utility.YamlToStruct("../../spec/app-spec.yaml", c)
	if err != nil {

	}
	p := DefaultPipeline{
		SOPath:   "/Users/anupam.gogoi.br/github/anupamgogoi/sonic/go-sonic-kafka-connector/plugin/",
		SonicApp: c,
	}
	p.ParseSonicApp()
}
