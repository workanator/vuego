package ui

import "gopkg.in/workanator/vuego.v1/html"

// ItemRenderer is responsible for rendering child components in parent bounds.
type ItemRenderer interface {
	html.Renderer

	// Items returns the slice of child components.
	Items() []Componenter

	// Len returns the number of child items.
	Len() int
}
