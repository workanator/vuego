package vuetify

import "gopkg.in/workanator/vuego.v1/html"

const (
	Default Appearance = iota
	Light
	Dark
)

type Appearance int

func (ap Appearance) IsDefault() bool {
	return ap == Default
}

func (ap Appearance) String() string {
	switch ap {
	case Light:
		return "light"
	case Dark:
		return "dark"
	default:
		return ""
	}
}

func (ap Appearance) Impose(el *html.Element) {
	if el != nil {
		switch ap {
		case Light, Dark:
			el.Attribute.Set(ap.String(), true)
		}
	}
}
