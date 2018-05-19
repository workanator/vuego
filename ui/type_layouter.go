package ui

import "gopkg.in/workanator/vuego.v1/html"

type Layouter interface {
	Layout(parent *html.Element, viewport html.Rect, items []Componenter) *html.Element
}
