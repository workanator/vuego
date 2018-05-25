package ui

import "gopkg.in/workanator/vuego.v1/html"

// Class of the Component.
type ComponentClass struct{}

func (ComponentClass) Class() string          { return "Component" }
func (ComponentClass) ExtendedClass() Classer { return nil }

// Component is the base class for all UI components.
type Component struct {
	Tag
}

// Get class name.
func (Component) Class() string {
	return ComponentClass{}.Class()
}

// Get the class name that class extends.
func (Component) ExtendedClass() Classer {
	return ComponentClass{}
}

// Render content into HTML Element.
func (Component) Render(parent *html.Element, viewport html.Rect) (*html.Element, error) {
	return nil, nil
}

// Impose attributes to HTML Element.
func (Component) Impose(el *html.Element) {}
