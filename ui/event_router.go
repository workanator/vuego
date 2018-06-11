package ui

import "gopkg.in/workanator/vuego.v1/session"

// Router is responsible for routing event inside component.
type EventRouter interface {
	RouteEvent(sess *session.Session, cmp Component, name string, data interface{}) error
}
