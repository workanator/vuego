package view

import "gopkg.in/workanator/vuego.v1/html"

type Screen struct {
	Id     string
	Name   string
	Title  string
	Head   html.Markuper
	Body   html.Renderer
	Models []html.Markuper
}
