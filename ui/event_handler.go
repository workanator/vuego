package ui

import (
	"gopkg.in/workanator/vuego.v1/action"
	"gopkg.in/workanator/vuego.v1/session"
)

// That is the direct handler of event. In other words that is the recipient of event.
type EventHandler interface {
	HandleEvent(sess *session.Session, cmp Component, data interface{}) error
}

// Wrap the standalone function into EventHandler and implement EventMarkuper interface
// to render proper event redirection from UI to back-end.
type Handle func(sess *session.Session, cmp Component, data interface{}) error

func (h Handle) HandleEvent(sess *session.Session, cmp Component, data interface{}) error {
	return h(sess, cmp, data)
}

func (h Handle) MarkupEvent(event Event, cmp Component) (string, error) {
	return "Vuego.Bus.send({category:'dom',target:'" + cmp.Target() + "',name:'" + event.String() + "',data:Vuego.Event.serialize(event)})", nil
}

// Wrap the standalone function into EventHandler and implement EventMarkuper interface
// to render proper action events.
type actionHandler action.Action

func Action(act *action.Action) *actionHandler {
	return (*actionHandler)(act)
}

func (*actionHandler) HandleEvent(sess *session.Session, cmp Component, data interface{}) error {
	return nil
}

func (a *actionHandler) MarkupEvent(event Event, cmp Component) (string, error) {
	return "Vuego.Action." + (*action.Action)(a).JS(), nil
}
