package ui

import "gopkg.in/workanator/vuego.v1/html"

type Component interface {
	Classer
	html.Renderer
	html.Imposer

	ProcessEvent(recipient, event string, data interface{}) (processed bool, err error)
}
