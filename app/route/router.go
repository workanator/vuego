package route

import (
	"gopkg.in/workanator/vuego.v1/app/session"
	"gopkg.in/workanator/vuego.v1/mvc/view"
)

type Router interface {
	Route(sess *session.Session, path string) (scr *view.Screen, err error)
}
