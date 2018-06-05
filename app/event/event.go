package event

import "time"

type Event struct {
	Source   EventSource   `json:"-"`
	Category EventCategory `json:"category"`
	Name     string        `json:"name"`
	Time     time.Time     `json:"time"`
	Data     interface{}   `json:"data,omitempty"`
}

// Event sources
const (
	SourceServer EventSource = "server"
	SourceClient EventSource = "client"
)

type EventSource string

// Test if the event source is server
func (es EventSource) IsServer() bool {
	return es == SourceServer
}

// Test if the event source is client
func (es EventSource) IsClient() bool {
	return es == SourceClient
}

// Test if the event source is custom
func (es EventSource) IsCustom() bool {
	return es != SourceServer && es != SourceClient
}

// Event categories
const (
	CategorySystem EventCategory = "system"
	CategoryModel  EventCategory = "model"
	CategoryEvent  EventCategory = "event"
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

// Test if the event category is event
func (ec EventCategory) IsEvent() bool {
	return ec == CategoryEvent
}

// Test if the event category is custom
func (ec EventCategory) IsCustom() bool {
	return ec != CategorySystem && ec != CategoryModel && ec != CategoryEvent
}
