package ui

import "gopkg.in/workanator/vuego.v1/html"

type Componenter interface {
	Classer
	html.Renderer
	html.Imposer
}
