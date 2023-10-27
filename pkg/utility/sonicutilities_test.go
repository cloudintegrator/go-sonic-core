package utility

import (
	"github.com/go-playground/assert"
	"testing"
)

type MockConnector struct {
	Kind          string `yaml:"kind"`
	Configuration struct {
		Host      string `yaml:"host"`
		Port      string `yaml:"port"`
		Topic     string `yaml:"topic"`
		Partition string `yaml:"partition"`
	}
}

func TestYamlToStruct(t *testing.T) {
	c := &MockConnector{}
	YamlToStruct("../../spec/mock-connector.yaml", c)
	assert.Equal(t, c.Configuration.Host, "localhost")
}
