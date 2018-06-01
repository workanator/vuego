package test

import (
	"gopkg.in/workanator/vuego.v1/event"
	"gopkg.in/workanator/vuego.v1/html"
	"gopkg.in/workanator/vuego.v1/model"
	"gopkg.in/workanator/vuego.v1/theme/vuetify"
	"gopkg.in/workanator/vuego.v1/ui"
	"gopkg.in/workanator/vuego.v1/ui/layout"
	"gopkg.in/workanator/vuego.v1/vue"
)

type App struct{}

// Get internal screen name.
func (a *App) Name() string {
	return "TestApp"
}

// Get user-friendly title.
func (a *App) Title() string {
	return "Test Application which does nothing"
}

// Get head tags.
func (a *App) Head() html.Markuper {
	return nil
}

// Get body renderer.
func (a *App) Body() html.Renderer {
	return &vuetify.App{
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
}

// Get models used on the screen.
func (a *App) Models() []html.Markuper {
	m1 := &model.Container{}
	m1.Field("message").Initial("Hello from test application!")

	m2 := &model.Container{}
	m2.Field("processed").Initial(false)

	return []html.Markuper{
		&vue.Vue{Id: "app", Data: m1},
		&vue.Vue{Id: "state", Data: m2},
	}
}

// Get screen's event producer.
func (a *App) EventProducer() event.Producer {
	return nil
}

// Get screen's event consumer.
func (a *App) EventConsumer() event.Consumer {
	return nil
}
