package test_todo

import (
	"gopkg.in/workanator/vuego.v1/app"
	"gopkg.in/workanator/vuego.v1/session"
	"gopkg.in/workanator/vuego.v1/theme/vuetify"
	"gopkg.in/workanator/vuego.v1/ui"
	"gopkg.in/workanator/vuego.v1/ui/layout"
)

type router struct{}

func (router) Screen(sess *session.Session, route string) (scr *app.Screen, err error) {
	body := &vuetify.App{
		Tag: ui.Tag{
			Id: "app",
		},
		Appearance: vuetify.Dark,
		Children: layout.Vert{
			&ui.Text{
				Text: "Application says '{{message}}'",
				Events: ui.Listeners{
					ui.OnClick: ui.HandlerFunc(func(sess *session.Session, cmp ui.Component, data interface{}) error {
						println("Clicked text")
						return nil
					}),
				},
			},
		},
	}

	return &app.Screen{
		Title: "To-Do List",
		Root:  body,
	}, nil
}
