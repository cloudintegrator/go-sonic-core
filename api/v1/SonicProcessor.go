package v1

type SonicProcessor interface {
	Process(p *SonicProcessor) SonicProcessor
}
