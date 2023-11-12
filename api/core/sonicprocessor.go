package core

// SonicProcessor Source is a special type of SonicComponent
type SonicProcessor interface {
	SonicComponent
	Process()
}
