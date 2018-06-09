package action

import (
	"gopkg.in/workanator/vuego.v1/session"
	"gopkg.in/workanator/vuego.v1/view"
)

type Responder interface {
	Respond(sess *session.Session, action string) (scr *view.Screen, err error)
}
