package core

import (
	"context"
	"sync"
)

// SonicContext The Context that will pass through all components (Source, Processor) of the Pipeline.
type SonicContext struct {
	WG   *sync.WaitGroup
	Ctx  *context.Context
	Data []interface{}
}
