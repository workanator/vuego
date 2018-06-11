package ui

import (
	"gopkg.in/workanator/vuego.v1/html"
	"gopkg.in/workanator/vuego.v1/session"
)

// Class of the El.
type ElClass struct{}

func (ElClass) Class() string          { return "El" }
func (ElClass) ExtendedClass() Classer { return nil }

// El is the base class for all UI components.
type El struct{}

// Get class name.
func (El) Class() string {
	return ElClass{}.Class()
}

// Get the class name that class extends.
func (El) ExtendedClass() Classer {
	return ElClass{}
}

// Render content into HTML Element.
func (El) Render(parent *html.Element, viewport html.Rect) (*html.Element, error) {
	return nil, nil
}

// Implement Component interface. Always return the event is not processed.
func (El) MarshalEvent(sess *session.Session, target, name string, data interface{}) (processed bool, err error) {
	return false, nil
}
