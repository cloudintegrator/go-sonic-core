package core

type SonicConnector interface {
	SonicConnectorConfig
	Say()
}

type SonicConnectorConfig struct {
}
