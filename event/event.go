package event

import "time"

type Event struct {
	Source   EventSource
	Category EventCategory
	Name     string
	Time     time.Time
	Data     interface{}
}

// Event sources
const (
	_ EventSource = iota
	SourceServer
	SourceClient
	SourceCustom
)

type EventSource uint8

// Event categories
const (
	_ EventCategory = iota
	CategoryData
	CategoryEvent
	CategoryCustom
)

type EventCategory uint8
