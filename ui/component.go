package ui

import (
	"gopkg.in/workanator/vuego.v1/html"
)

type Component interface {
	Classer
	html.Renderer
	EventMarshaler
}
