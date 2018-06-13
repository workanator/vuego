package ui

import (
	"gopkg.in/workanator/vuego.v1/html"
	"gopkg.in/workanator/vuego.v1/session"
)

// Set of event listeners which are responsible for event handling.
type Listeners map[Event]EventHandler

func (eh Listeners) RouteEvent(sess *session.Session, cmp Component, name string, data interface{}) error {
	if eh == nil {
		return ErrNil{}
	}

	// Call the handler if exists
	if handler := eh[Event(name)]; handler != nil {
		return handler.HandleEvent(sess, cmp, data)
	}

	return nil
}

func (eh Listeners) ImposeEvent(el *html.Element, cmp Component) error {
	if eh == nil {
		return nil
	}

	for name := range eh {
		if handler := eh[Event(name)]; handler != nil {
			switch v := handler.(type) {
			case EventMarkuper:
				if markup, err := v.MarkupEvent(name, cmp); err != nil {
					return err
				} else {
					el.Attribute.Set(name.String(), markup)
				}

			case html.Markuper:
				if markup, err := v.Markup(); err != nil {
					return err
				} else {
					el.Attribute.Set(name.String(), markup)
				}

			case EventImposer:
				if err := v.ImposeEvent(el, cmp); err != nil {
					return err
				}

			case html.Imposer:
				v.Impose(el)
			}
		}
	}

	return nil
}
