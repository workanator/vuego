package ui

import "gopkg.in/workanator/vuego.v1/html"

type Bounds struct {
	Rect     html.Rect
	Position html.Position
	Overflow html.Overflow
}

func (b *Bounds) Impose(el *html.Element) *html.Element {
	if el != nil {
		b.Rect.Impose(el)
		b.Position.Impose(el)
		b.Overflow.Impose(el)
	}

	return el
}
