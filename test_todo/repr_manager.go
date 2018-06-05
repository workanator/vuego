package test_todo

import (
	"gopkg.in/workanator/vuego.v1/app"
	"gopkg.in/workanator/vuego.v1/app/session"
	"gopkg.in/workanator/vuego.v1/html"
	"gopkg.in/workanator/vuego.v1/mvc/model"
	"gopkg.in/workanator/vuego.v1/theme/vuetify"
	"gopkg.in/workanator/vuego.v1/ui"
	"gopkg.in/workanator/vuego.v1/ui/layout"
	"gopkg.in/workanator/vuego.v1/vue"
)

type reprManager struct{}

func (reprManager) Representation(sess *session.Session) (scr *app.Screen, err error) {
	body := &vuetify.App{
		Tag: ui.Tag{
			Id: "app",
		},
		Appearance: vuetify.Dark,
		Children: layout.Vert{
			&ui.Text{
				Text: "Application says '{{message}}'",
			},
		},
	}

	m1 := &model.Container{}
	m1.Field("message").Initial("Hello from test application!")

	m2 := &model.Container{}
	m2.Field("processed").Initial(false)

	models := []html.Markuper{
		&vue.Vue{Id: "app", Data: m1},
		&vue.Vue{Id: "state", Data: m2},
	}

	return &app.Screen{
		Id:     "list",
		Name:   "list",
		Title:  "To-Do List",
		Head:   nil,
		Body:   body,
		Models: models,
	}, nil
}
