package core

// Source is a special type of SonicCompo
type SonicSource interface {
	SonicComponent
	Listen() string
}
