package ui

import "gopkg.in/workanator/vuego.v1/html"

type Imposer interface {
	Impose(el *html.Element)
}
