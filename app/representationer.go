package app

import "gopkg.in/workanator/vuego.v1/app/session"

type Representationer interface {
	Representation(sess *session.Session) (scr *Screen, err error)
}
