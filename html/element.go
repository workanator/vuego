package html

import "strings"

type Element struct {
	Tag       string
	Id        Ider
	Class     Classer
	Style     Styler
	Attribute Attributer
	Inner     Markuper
}

func (el *Element) Markup() string {
	markup := strings.Builder{}
	markup.WriteRune('<')
	markup.WriteString(el.Tag)

	if el.Id != nil {
		markup.WriteString(el.Id.Markup())
	}

	if el.Class != nil {
		markup.WriteString(el.Class.Markup())
	}

	if el.Style != nil {
		markup.WriteString(el.Style.Markup())
	}

	if el.Attribute != nil {
		markup.WriteString(el.Attribute.Markup())
	}

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
