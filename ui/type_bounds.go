package ui

import "gopkg.in/workanator/vuego.v1/html"

type Bounds struct {
	Rect
	Position Position
	Overflow Overflow
}

func (b *Bounds) Impose(el *html.Element) {
	if el != nil {
		// Impose bounding rect
		if b.Rect.HasTop() {
			el.Style.Setf("top", "%dpx", b.Rect.Top)
		}

		if b.Rect.HasBottom() {
			el.Style.Setf("bottom", "%dpx", b.Rect.Bottom)
		}

		if b.Rect.HasLeft() {
			el.Style.Setf("left", "%dpx", b.Rect.Left)
		}

		if b.Rect.HasRight() {
			el.Style.Setf("right", "%dpx", b.Rect.Right)
		}

		// Impose positioning and overflow style
		b.Position.Impose(el)
		b.Overflow.Impose(el)
	}
}
