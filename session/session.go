package session

import (
	"context"

	"gopkg.in/workanator/vuego.v1/event"
)

type Session struct {
	Context  context.Context
	Id       string
	User     *User
	Data     interface{}
	EventBus event.Bus
}
