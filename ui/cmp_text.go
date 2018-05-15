package ui

import "gopkg.in/workanator/vuego.v1/html"

// Class of the Text.
type TextClass struct{}

func (TextClass) Class() string          { return "Text" }
func (TextClass) ExtendedClass() Classer { return ComponentClass{} }

// Type of the text container.
const (
	TextPlain TextType = iota
	TextParagraph
	TextCite
	TextBlockquote
	TextPreformatted
	TextCode
)

type TextType uint8

// Simple text component.
type Text struct {
	Tag
	Bounds
	Text string
	Type TextType
}

// Get class name.
func (Text) Class() string {
	return TextClass{}.Class()
}

// Get the class name that class extends.
func (Text) ExtendedClass() Classer {
	return TextClass{}
}

// Render content into HTML Element.
func (txt *Text) Render(parent *html.Element, viewport Rect) *html.Element {
	el := txt.Tag.Element()
	el.Class.Add("vg-text")

	switch txt.Type {
	case TextPlain:
		el.Tag = "span"
		el.Inner = html.Text(txt.Text)
	case TextParagraph:
		el.Tag = "p"
		el.Inner = html.Text(txt.Text)
	case TextCite:
		el.Tag = "cite"
		el.Inner = &html.Element{
			Tag:   "p",
			Inner: html.Text(txt.Text),
		}
	case TextBlockquote:
		el.Tag = "blockquote"
		el.Inner = &html.Element{
			Tag:   "p",
			Inner: html.Text(txt.Text),
		}
	case TextPreformatted:
		el.Tag = "pre"
		el.Inner = html.Text(txt.Text)
	case TextCode:
		el.Tag = "code"
		el.Inner = html.Text(txt.Text)
	}

	txt.Bounds.Impose(el)

	return el
}
