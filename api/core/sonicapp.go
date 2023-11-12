package core

// SonicApp is the data model of the yaml application configuration.
// When the yaml application is read by the ESB it's converted to this structure.
type SonicApp struct {
	App App `yaml:"app"`
}

type App struct {
	Name   string   `yaml:"name"`
	Kind   string   `yaml:"kind"`
	Common []Common `yaml:"common"`
	Flows  []Flow   `yaml:"flows"`
}

type Common struct {
	Name          string      `yaml:"name"`
	Kind          string      `yaml:"kind"`
	Configuration interface{} `yaml:"configuration"`
}

type Flow struct {
	Name       string      `yaml:"name"`
	Kind       string      `yaml:"kind"`
	Components []Component `yaml:"components"`
}

type Component struct {
	Name          string      `yaml:"name"`
	Kind          string      `yaml:"kind"`
	Configuration interface{} `yaml:"configuration,omitempty"`
}
