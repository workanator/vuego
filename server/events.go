package server

import (
	"gopkg.in/workanator/vuego.v1/event"
)

const (
	targetBus = "ws.bus"
)

const (
	nameWsClosed = "closed_by_client"
)

var (
	evClientClosedWebSocket = event.Event{
		Category: event.CategorySystem,
		Target:   targetBus,
		Name:     nameWsClosed,
	}
)
