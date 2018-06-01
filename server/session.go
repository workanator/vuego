package server

import (
	"context"

	"gopkg.in/workanator/vuego.v1/app"
	"gopkg.in/workanator/vuego.v1/event"
	"gopkg.in/workanator/vuego.v1/session"
)

type Session struct {
	User      *session.User
	Screens   []app.Screener
	ServerBus *event.Bus
	ClientBus *event.Bus
	Context   context.Context
}

func NewSession(user *session.User, screen app.Screener) (*Session, error) {
	return &Session{
		User:      user,
		Screens:   []app.Screener{screen},
		ServerBus: event.NewBus(),
		ClientBus: event.NewBus(),
		Context:   context.Background(),
	}, nil
}
