package app

import "gopkg.in/workanator/vuego.v1/session"

type Screener interface {
	Screen(sess *session.Session, route string) (scr *Screen, err error)
}
