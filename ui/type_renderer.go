package ui

import (
	"gopkg.in/workanator/vuego.v1/html"
)

type Renderer interface {
	// Render content into HTML Element.
	Render(parent *html.Element, viewport Rect) *html.Element
}
