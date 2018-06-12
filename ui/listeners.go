package ui

import "gopkg.in/workanator/vuego.v1/session"

// Set of event listeners which are responsible for event handling.
type Listeners map[Event]EventHandler

func (eh Listeners) RouteEvent(sess *session.Session, cmp Component, name string, data interface{}) error {
	if eh == nil {
		return ErrNil{}
	}

	// Call the handler if exists
	if handler := eh[Event(name)]; handler != nil {
		return handler(sess, cmp, data)
	}

	return nil
}
