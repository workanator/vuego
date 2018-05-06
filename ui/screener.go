package ui

type Screener interface {
	Renderer

	GetName() string
	GetTitle() string
}
