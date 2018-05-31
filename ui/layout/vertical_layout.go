package layout

import (
	"fmt"

	"gopkg.in/workanator/vuego.v1/html"
	"gopkg.in/workanator/vuego.v1/ui"
)

type Vert []ui.Componenter

func (l Vert) Render(parent *html.Element, viewport html.Rect) (*html.Element, error) {
	// Return nil if there are no children.
	if len(l) == 0 {
		return nil, nil
	}

	// Render children items in put them into the element.
	items := make([]html.Markuper, len(l))
	for k, item := range l {
		if el, err := item.Render(parent, viewport); err != nil {
			return nil, html.ErrRenderFailed{
				Class:  item.Class(),
				Id:     fmt.Sprintf("vertical_layout_item_%d", k),
				Reason: err,
			}
		} else {
			items[k] = el
		}
	}

	return &html.Element{
		Tag:   "v-container",
		Inner: html.Multiple(items),
	}, nil
}

func (l Vert) Items() []ui.Componenter {
	return l
}

func (l Vert) Len() int {
	return len(l)
}
