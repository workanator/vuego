package event

import "time"

type Event struct {
	Target   string        `json:"target"`
	Category EventCategory `json:"category"`
	Name     string        `json:"name"`
	Time     time.Time     `json:"time"`
	Data     interface{}   `json:"data,omitempty"`
}

func (e *Event) Conforms(target, name string) bool {
	return e.Target == target && e.Name == name
}

// Event categories
const (
	CategorySystem EventCategory = "system"
	CategoryModel  EventCategory = "model"
	CategoryDom    EventCategory = "dom"
	CategoryUser   EventCategory = "user"
)

type EventCategory string

// Test if the event category is system
func (ec EventCategory) IsSystem() bool {
	return ec == CategorySystem
}

// Test if the event category is model
func (ec EventCategory) IsModel() bool {
	return ec == CategoryModel
}

// Test if the event category is DOM
func (ec EventCategory) IsDom() bool {
	return ec == CategoryDom
}

// Test if the event category is user
func (ec EventCategory) IsUser() bool {
	return ec == CategoryUser
}

// Test if the event category is custom
func (ec EventCategory) IsCustom() bool {
	return ec != CategorySystem && ec != CategoryModel && ec != CategoryDom && ec != CategoryUser
}
