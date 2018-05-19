package vuetify

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
