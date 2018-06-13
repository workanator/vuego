package ui

import "gopkg.in/workanator/vuego.v1/html"

type EventImposer interface {
	ImposeEvent(el *html.Element, cmp Component) error
}
