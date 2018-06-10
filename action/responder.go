package action

import (
	"gopkg.in/workanator/vuego.v1/session"
)

type Responder interface {
	Respond(sess *session.Session, action string) (resp interface{}, err error)
}
