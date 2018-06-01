package server

import (
	"context"

	"gopkg.in/workanator/vuego.v1/app"
	"gopkg.in/workanator/vuego.v1/event"
	"gopkg.in/workanator/vuego.v1/session"
)

type Session struct {
	Context        context.Context
	User           *session.User
	Screens        []app.Screener
	OutboundEvents writeBusConfig
	InboundEvents  readBusConfig
}

type writeBusConfig struct {
	Bus    *event.Bus
	Pusher event.ProducePusher
}

type readBusConfig struct {
	Bus    *event.Bus
	Puller event.ConsumePuller
}

func (server *Server) newSession(user *session.User, screen app.Screener) (*Session, error) {
	pusher := event.NewPushQueue(server.OutboundQueueSize)
	puller := event.NewPullQueue(server.InboundQueueSize)

	return &Session{
		Context: context.Background(),
		User:    user,
		Screens: []app.Screener{screen},
		OutboundEvents: writeBusConfig{
			Bus:    event.NewBus(),
			Pusher: &pusher,
		},
		InboundEvents: readBusConfig{
			Bus:    event.NewBus(),
			Puller: &puller,
		},
	}, nil
}
