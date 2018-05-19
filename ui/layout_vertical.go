package ui

import "gopkg.in/workanator/vuego.v1/html"

type LayoutVertical []Componenter

func (l LayoutVertical) Layout(parent *html.Element, viewport html.Rect) *html.Element {
	// Return nil if there are no children.
	if len(l) == 0 {
		return nil
	}

	// Render children items in put them into the element.
	items := make([]html.Markuper, len(l))
	for k, item := range l {
		items[k] = item.Render(parent, viewport)
	}

	return &html.Element{
		Tag:   "v-container",
		Inner: html.Multiple(items),
	}
}

func (l LayoutVertical) Items() []Componenter {
	return l
}
