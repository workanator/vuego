package view

import "gopkg.in/workanator/vuego.v1/session"

type Finder interface {
	Find(sess *session.Session, route string) (view *View, err error)
}
