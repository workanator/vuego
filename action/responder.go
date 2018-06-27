package action

import (
	"gopkg.in/workanator/vuego.v1/session"
)

type Responder interface {
	Respond(sess *session.Session, act *Action) (resp interface{}, err error)
}
