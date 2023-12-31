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
	FindComponentInSO("/Users/anupam.gogoi.br/github/anupamgogoi/sonic/go-sonic-kafka-connector/plugin/", "EXPORT_KAFKA_LISTENER", "kafka:listener")
}

func TestGetSOPath(t *testing.T) {
	GetSOPath("/Users/anupam.gogoi.br/github/anupamgogoi/sonic/go-sonic-kafka-connector/plugin/", "kafka:listener")
}
