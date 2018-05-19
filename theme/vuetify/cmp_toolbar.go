package vuetify

import (
	"gopkg.in/workanator/vuego.v1/html"
	"gopkg.in/workanator/vuego.v1/ui"
)

// Class of the Toolbar.
type ToolbarClass struct{}

func (ToolbarClass) Class() string             { return "Toolbar" }
func (ToolbarClass) ExtendedClass() ui.Classer { return ui.ComponentClass{} }

// Vuetify Toolbar.
type Toolbar struct {
	ui.Tag
	Layout          ui.Layouter
	Appearance      Appearance
	Absolute        bool
	Fixed           bool
	App             bool
	Card            bool
	ClippedLeft     bool
	ClippedRight    bool
	Dense           bool
	Extended        bool
	ExtensionHeight string
	Flat            bool
	Floating        bool
	Height          string
	Prominent       bool
	Tabs            bool
	Scroll          struct {
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
func (tb *Toolbar) Render(parent *html.Element, viewport html.Rect) *html.Element {
	if len(tb.Tag.Id) == 0 {
		tb.Tag.Id = "app"
	}

	el := tb.Tag.Element()
	el.Tag = "v-toolbar"

	return el
}
