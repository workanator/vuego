package test_todo

import (
	"gopkg.in/workanator/vuego.v1/app"
	"gopkg.in/workanator/vuego.v1/event"
	"gopkg.in/workanator/vuego.v1/facade"
)

func Bundle() app.Bundle {
	pushQue := event.NewPushQueue(64)
	pullQue := event.NewPullQueue(64)

	return app.Bundle{
		Id:      "todo",
		Name:    "Simple To-Do",
		Version: "1.0",
		Fs:      &facade.FileSystem{},
		Sessions: &facade.SingleSession{
			InboundEvents:  &pushQue,
			OutboundEvents: &pullQue,
		},
		Screens: &reprManager{},
	}
}
