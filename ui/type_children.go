package ui

import "gopkg.in/workanator/vuego.v1/html"

type Children struct {
	Layout string
	Items  []Componenter
}

func (ch *Children) Len() int {
	return len(ch.Items)
}

func (ch *Children) Impose(el *html.Element) {
}

func (ch *Children) Render(parent *html.Element, viewport Rect) *html.Element {
	// Return nil if there are no children.
	if len(ch.Items) == 0 {
		return nil
	}

	// Make the list of markapers in put them into the element.
	items := make([]html.Markuper, len(ch.Items))
	for k, item := range ch.Items {
		items[k] = item.Render(parent, viewport)
	}

	return &html.Element{
		Inner: html.Multiple(items),
	}
}
