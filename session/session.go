package session

import (
	"context"

	"gopkg.in/workanator/vuego.v1/event"
)

type Session struct {
	Context  context.Context
	Id       string
	User     *User
	State    interface{}
	Inbound  event.Consumer
	Outbound event.Producer
}
