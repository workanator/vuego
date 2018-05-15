package ui

import "gopkg.in/workanator/vuego.v1/html"

type Tag struct {
	Id        html.Id
	Class     html.Class
	Style     html.Style
	Attribute html.Attribute
}

func (t *Tag) Element() *html.Element {
	return &html.Element{
		Id:        t.Id,
		Class:     t.Class,
		Style:     t.Style,
		Attribute: t.Attribute,
	}
}
