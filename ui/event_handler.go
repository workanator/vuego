package ui

// That is the direct handler of event. In other words that is the recipient of event.
type EventHandler interface {
	HandleEvent(cmp Component, data interface{}) error
}

// Wrapper for event handling function.
type HandlerFunc func(Component, interface{}) error

func (hf HandlerFunc) HandleEvent(cmp Component, data interface{}) error {
	return hf(cmp, data)
}
