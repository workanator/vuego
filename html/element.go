package html

import (
	"strings"

	"gopkg.in/workanator/vuego.v1/errors"
)

type Element struct {
	Tag       string
	Short     bool
	Id        Id
	Class     Class
	Style     Style
	Attribute Attribute
	Inner     Markuper
}

func (el *Element) Markup() (string, error) {
	// Render inner tag
	var inner string
	if el.Inner != nil {
		var err error
		if inner, err = el.Inner.Markup(); err != nil {
			return "", errors.ErrMarkupFailed{
				Tag:    el.Tag,
				Id:     el.Id.String(),
				Reason: err,
			}
		}
	}

	// Render Inner elements only if tag is empty
	if len(el.Tag) == 0 {
		return inner, nil
	}

	// Render the whole element
	markup := strings.Builder{}
	markup.WriteRune('<')
	markup.WriteString(el.Tag)

	// Render markup for id attribute
	if id, err := el.Id.Markup(); err != nil {
		return "", errors.ErrMarkupFailed{
			Tag:    el.Tag,
			Id:     el.Id.String(),
			Reason: err,
		}
	} else {
		markup.WriteString(id)
	}

	// Render markup for class attribute
	if class, err := el.Class.Markup(); err != nil {
		return "", errors.ErrMarkupFailed{
			Tag:    el.Tag,
			Id:     el.Id.String(),
			Reason: err,
		}
	} else {
		markup.WriteString(class)
	}

	// Render markup for style attribute
	if style, err := el.Style.Markup(); err != nil {
		return "", errors.ErrMarkupFailed{
			Tag:    el.Tag,
			Id:     el.Id.String(),
			Reason: err,
		}
	} else {
		markup.WriteString(style)
	}

	// Render markup for other attributes
	if attrs, err := el.Attribute.Markup(); err != nil {
		return "", errors.ErrMarkupFailed{
			Tag:    el.Tag,
			Id:     el.Id.String(),
			Reason: err,
		}
	} else {
		markup.WriteString(attrs)
	}

	if el.Inner == nil {
		if el.Short {
			markup.WriteString("/")
		} else {
			markup.WriteRune('>')
			markup.WriteString("</")
			markup.WriteString(el.Tag)
		}
	} else {
		markup.WriteRune('>')
		markup.WriteString(inner)
		markup.WriteString("</")
		markup.WriteString(el.Tag)
	}

	markup.WriteRune('>')

	return markup.String(), nil
}
