package ui

type Traverser interface {
	Up(query string) Traverser

	Down(query string) Traverser

	Sibling(query string) Traverser

	Get() []Container

	GetOne() Container
}
