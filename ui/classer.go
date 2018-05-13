package ui

// Classer is the class of the component.
type Classer interface {
	// Get class name.
	Name() string

	// Extends returns the list of class that class extends.
	Extends() []Classer

	// Is tests if the class or any of extended classes is of cls class.
	Is(cls string) bool
}
