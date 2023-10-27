package core

type SonicProcessor interface {
	Process(p *SonicProcessor) SonicProcessor
}
