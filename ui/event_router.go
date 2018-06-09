package ui

// Router is responsible for routing event inside component.
type EventRouter interface {
	RouteEvent(cmp Component, name string, data interface{}) error
}
