package ui

import (
	"gopkg.in/workanator/vuego.v1/html"
	"gopkg.in/workanator/vuego.v1/session"
)

// Class of the Html.
type HtmlClass struct{}

func (HtmlClass) Class() string          { return "Html" }
func (HtmlClass) ExtendedClass() Classer { return ElClass{} }

// Custom HTML.
type Html struct {
	Tag
	Bounds
	Content string
	Events  EventRouter
}

// Get class name.
func (Html) Class() string {
	return HtmlClass{}.Class()
}

// Get the class name that class extends.
func (Html) ExtendedClass() Classer {
	return HtmlClass{}
}

// Render content into HTML Element.
func (cmp *Html) Render(parent *html.Element, viewport html.Rect) (*html.Element, error) {
	el := cmp.Tag.Element()
	el.Id = "test"
	el.Tag = "div"
	el.Class.Add("vg-html")
	el.Inner = html.Html(cmp.Content)

	cmp.Bounds.Impose(el)
	ApplyToEl(el, cmp, cmp.Events)

	return el, nil
}

// Implement Component interface.
func (cmp *Html) MarshalEvent(sess *session.Session, target, name string, data interface{}) (processed bool, err error) {
	if cmp.Tag.Id.Equal(target) {
		processed, err = true, cmp.Events.RouteEvent(sess, cmp, name, data)
	}

	return processed, err
}
