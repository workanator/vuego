package test

import (
	"gopkg.in/workanator/vuego.v1/app"
	"gopkg.in/workanator/vuego.v1/html"
	"gopkg.in/workanator/vuego.v1/model"
	"gopkg.in/workanator/vuego.v1/theme/vuetify"
	"gopkg.in/workanator/vuego.v1/ui"
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
		Children: ui.VerticalLayout{
			&ui.Text{
				Text: "Application says '{{message}}'",
			},
		},
	}
}

// Get models used on the screen.
func (a *App) Models() []app.ModelMarkuper {
	m1 := &model.SharedModel{
		Id: "app",
	}
	m1.Field("message").Initial("Hello from test application!")

	m2 := &model.SharedModel{
		Id: "state",
	}
	m2.Field("processed").Initial(false)

	return []app.ModelMarkuper{m1, m2}
}
