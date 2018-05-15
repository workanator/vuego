package ui

type Classer interface {
	// Get class name.
	Class() string

	// Get the class name that class extends. The return value is nil if no class is exetnded.
	ExtendedClass() Classer
}
