package session

import (
	"context"

	"gopkg.in/workanator/vuego.v1/event"
)

type Session struct {
	Context        context.Context
	Id             string
	User           *User
	State          interface{}
	InboundEvents  event.Pusher // Must be event.Producer
	OutboundEvents event.Puller // Must be event.Consumer
}
