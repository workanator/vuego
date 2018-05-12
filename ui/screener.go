package ui

type Screener interface {
	Renderer

	Name() string
	Title() string
}
