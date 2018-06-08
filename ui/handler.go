package ui

type Handler interface {
	Handle(cmp Component, data interface{}) error
}

type HandlerFunc func(Component, interface{}) error

func (hf HandlerFunc) Handle(cmp Component, data interface{}) error {
	return hf(cmp, data)
}

type EventHandler interface {
	HandleEvent(cmp Component, name string, data interface{}) error
}

type Listeners map[string]Handler

func (eh Listeners) HandleEvent(cmp Component, name string, data interface{}) error {
	if eh == nil {
		return ErrNil{}
	}

	// Call the handler if exists
	if handler := eh[name]; handler != nil {
		return handler.Handle(cmp, data)
	}

	return nil
}
