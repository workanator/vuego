package ui

type Container interface {
	// Get the identifier of the container.
	Id() string

	// Get the class of the container.
	Class() Classer

	// Get the parent container which holds that container. The parent returned may be nil when the container
	// has no parent.
	Parent() Container

	// Test if the container has parent container.
	HasParent() bool

	// Get the list of children containers the container is parent of. The list returned can be nil.
	Children() []Container

	// Test if the container has children containers.
	HasChildren() bool

	// Get the traversing facility.
	Traverse() Traverser
}
