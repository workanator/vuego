package server

import (
	"time"

	"gopkg.in/workanator/vuego.v1/app/event"
)

const (
	targetBus = "ws.bus"
)

const (
	nameWsClosed = "closed_by_client"
)

var (
	evClientClosedWebSocket = event.Event{
		Target:   targetBus,
		Category: event.CategorySystem,
		Name:     nameWsClosed,
		Time:     time.Now(),
	}
)
