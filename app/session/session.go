package session

import (
	"context"

	"gopkg.in/workanator/vuego.v1/app/event"
)

type Session struct {
	Context  context.Context
	Id       string
	User     *User
	State    interface{}
	EventBus event.Bus
}
