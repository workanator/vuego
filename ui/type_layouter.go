package ui

import "gopkg.in/workanator/vuego.v1/html"

// Layouter is responsible for laying out child components in parent bounds.
type Layouter interface {
	// Layout renders child components in the parent. The value returned can be nil
	// then layout contains no children.
	Layout(parent *html.Element, viewport html.Rect) *html.Element

	// Items returns the slice of child components.
	Items() []Componenter
}
