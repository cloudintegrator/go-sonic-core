package utility

import (
	"github.com/cloudintegrator/go-sonic-core/api/core"
	"log/slog"
	"testing"
)

func TestYamlToStruct(t *testing.T) {
	c := &core.SonicApp{}
	YamlToStruct("../../spec/app-spec.yaml", c)
	slog.Debug(c.App.Name)
}

func TestFindComponentInSO(t *testing.T) {
	FindComponentInSO("", "kafka:listener")
}

func TestGetSOPath(t *testing.T) {
	GetSOPath("kafka:listener")
}
