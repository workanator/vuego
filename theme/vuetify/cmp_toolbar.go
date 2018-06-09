package vuetify

import (
	"gopkg.in/workanator/vuego.v1/html"
	"gopkg.in/workanator/vuego.v1/ui"
)

// Class of the Toolbar.
type ToolbarClass struct{}

func (ToolbarClass) Class() string             { return "Toolbar" }
func (ToolbarClass) ExtendedClass() ui.Classer { return ui.ElClass{} }

// Vuetify Toolbar.
type Toolbar struct {
	ui.Tag
	Layout          ui.ItemRenderer
	Appearance      Appearance
	Position        html.Position // Absolute and Fixed
	Card            bool
	Dense           bool
	Extended        bool
	ExtensionHeight string
	Flat            bool
	Floating        bool
	Height          string
	Prominent       bool
	Tabs            bool
	Clipped         struct {
		Left  bool
		Right bool
	}
	Scroll struct {
		Inverted         bool
		Manual           bool
		OffScreen        bool
		Target           string
		Threshold        int
		ToolbarOffScreen bool
	}
}

// Get class name.
func (Toolbar) Class() string {
	return ToolbarClass{}.Class()
}

// Get the class name that class extends.
func (Toolbar) ExtendedClass() ui.Classer {
	return ToolbarClass{}
}

// Render content into HTML Element.
func (tb *Toolbar) Render(parent *html.Element, viewport html.Rect) (*html.Element, error) {
	el := tb.Tag.Element()
	el.Tag = "v-toolbar"
	tb.Impose(el)

	return el, nil
}

// Apply attributes to HTML Element.
func (tb *Toolbar) Impose(el *html.Element) {
	if el != nil {
		tb.Appearance.Impose(el)

		switch tb.Position {
		case html.PositionAbsolute:
			el.Attribute.Set("absolute", true)
		case html.PositionFixed:
			el.Attribute.Set("fixed", true)
		}

		if tb.Card {
			el.Attribute.Set("card", true)
		}

		if tb.Dense {
			el.Attribute.Set("dense", true)
		}

		if tb.Extended {
			el.Attribute.Set("extended", true)
		}

		if len(tb.ExtensionHeight) > 0 {
			el.Attribute.Set("extension-height", tb.ExtensionHeight)
		}

		if tb.Flat {
			el.Attribute.Set("flat", true)
		}

		if tb.Floating {
			el.Attribute.Set("floating", tb.Floating)
		}

		if len(tb.Height) > 0 {
			el.Attribute.Set("height", tb.Height)
		}

		if tb.Prominent {
			el.Attribute.Set("prominent", true)
		}

		if tb.Tabs {
			el.Attribute.Set("tabs", true)
		}
	}
}

// Implement Component interface.
func (tb *Toolbar) Process(recipient, event string, data interface{}) (processed bool, err error) {
	return processed, err
}
