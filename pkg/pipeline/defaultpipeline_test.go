package pipeline

import (
	"github.com/cloudintegrator/go-sonic-core/api/core"
	"github.com/cloudintegrator/go-sonic-core/pkg/utility"
	"testing"
)

func TestBuildPipeline(t *testing.T) {
	c := &core.SonicApp{}
	utility.YamlToStruct("../../spec/app-spec.yaml", c)
	ParseSonicApp(c)
}
