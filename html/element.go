package html

import "strings"

type Element struct {
	Tag       string
	Id        Id
	Class     Class
	Style     Style
	Attribute Attribute
	Inner     Markuper
}

func (el *Element) Markup() string {
	if len(el.Tag) == 0 {
		// Render Inner elements only
		if el.Inner != nil {
			return el.Inner.Markup()
		}

		// Render nothing
		return ""
	}

	// Render the whole element
	markup := strings.Builder{}
	markup.WriteRune('<')
	markup.WriteString(el.Tag)
	markup.WriteString(el.Id.Markup())
	markup.WriteString(el.Class.Markup())
	markup.WriteString(el.Style.Markup())
	markup.WriteString(el.Attribute.Markup())

	if el.Inner == nil {
		markup.WriteString("/")
	} else {
		markup.WriteRune('>')
		markup.WriteString(el.Inner.Markup())
		markup.WriteString("</")
		markup.WriteString(el.Tag)
	}

	markup.WriteRune('>')

	return markup.String()
}
