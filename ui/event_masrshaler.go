package ui

import "gopkg.in/workanator/vuego.v1/session"

// Marshaler delivers event to the proper recipient in the component tree.
type EventMarshaler interface {
	MarshalEvent(sess *session.Session, target, name string, data interface{}) (processed bool, err error)
}
