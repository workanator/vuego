package ui

import "gopkg.in/workanator/vuego.v1/session"

// That is the direct handler of event. In other words that is the recipient of event.
type EventHandler interface {
	HandleEvent(sess *session.Session, cmp Component, data interface{}) error
}

// Wrapper for event handling function.
type HandlerFunc func(*session.Session, Component, interface{}) error

func (hf HandlerFunc) HandleEvent(sess *session.Session, cmp Component, data interface{}) error {
	return hf(sess, cmp, data)
}
