package view

import (
	"gopkg.in/workanator/vuego.v1/html"
	"gopkg.in/workanator/vuego.v1/ui"
)

type Screen struct {
	Id     string
	Name   string
	Title  string
	Head   html.Markuper
	Body   ui.Component
	Models []html.Markuper
}
