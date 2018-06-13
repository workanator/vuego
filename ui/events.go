package ui

const (
	OnClick Event = "onclick"
)

// The predefined set of possible events running in UI.
type Event string

func (e Event) String() string {
	return string(e)
}
