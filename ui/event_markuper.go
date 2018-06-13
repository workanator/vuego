package ui

type EventMarkuper interface {
	MarkupEvent(event Event, cmp Component) (string, error)
}
