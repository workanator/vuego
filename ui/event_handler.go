package ui

import "gopkg.in/workanator/vuego.v1/session"

// That is the direct handler of event. In other words that is the recipient of event.
type EventHandler func(sess *session.Session, cmp Component, data interface{}) error
