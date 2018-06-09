package ui

// Set of event listeners which are responsible for event handling.
type Listeners map[Event]EventHandler

func (eh Listeners) HandleEvent(cmp Component, name string, data interface{}) error {
	if eh == nil {
		return ErrNil{}
	}

	// Call the handler if exists
	if handler := eh[Event(name)]; handler != nil {
		return handler.HandleEvent(cmp, data)
	}

	return nil
}
