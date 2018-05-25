package app

import (
	"gopkg.in/workanator/vuego.v1/html"
	"gopkg.in/workanator/vuego.v1/model"
)

type ModelMarkuper interface {
	model.Modeler
	html.Markuper
}
