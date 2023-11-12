package context

import (
	"context"
	"github.com/cloudintegrator/go-sonic-core/api/core"
	"sync"
)

type SonicContextBuilder interface {
	WG(WG *sync.WaitGroup) SonicContextBuilder
	Ctx(Ctx *context.Context) SonicContextBuilder
	Data(Data interface{}) SonicContextBuilder
	Build() *core.SonicContext
}

type sonicContextBuilder struct {
	wg   *sync.WaitGroup
	ctx  *context.Context
	data interface{}
}

func (s *sonicContextBuilder) WG(WG *sync.WaitGroup) SonicContextBuilder {
	s.wg = WG
	return s
}
func (s *sonicContextBuilder) Ctx(Ctx *context.Context) SonicContextBuilder {
	s.ctx = Ctx
	return s
}

func (s *sonicContextBuilder) Data(Data interface{}) SonicContextBuilder {
	s.data = Data
	return s
}
func (s *sonicContextBuilder) Build() *core.SonicContext {
	sc := &core.SonicContext{
		WG:   s.wg,
		Ctx:  s.ctx,
		Data: s.data,
	}
	return sc
}
