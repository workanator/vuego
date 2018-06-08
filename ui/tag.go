package ui

import (
	"fmt"
	"sync/atomic"

	"gopkg.in/workanator/vuego.v1/html"
)

var nextTagId uint64

type Tag struct {
	Id        html.Id
	Class     html.Class
	Style     html.Style
	Attribute html.Attribute
}

// Generate Id if it is empty.
func (t *Tag) CheckId() {
	if t.Id.IsEmpty() {
		t.Id = html.Id(fmt.Sprintf("A%06d", atomic.AddUint64(&nextTagId, 1)))
	}
}

func (t *Tag) Element() *html.Element {
	t.CheckId()

	return &html.Element{
		Id:        t.Id,
		Class:     t.Class,
		Style:     t.Style,
		Attribute: t.Attribute,
	}
}
