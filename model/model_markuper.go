package model

import "gopkg.in/workanator/vuego.v1/html"

type ModelMarkuper interface {
	Modeler
	html.Markuper
}
