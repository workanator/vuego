package todo

import (
	"gopkg.in/workanator/vuego.v1/session"
	"gopkg.in/workanator/vuego.v1/theme/vuetify"
	"gopkg.in/workanator/vuego.v1/ui"
	"gopkg.in/workanator/vuego.v1/ui/layout"
	"gopkg.in/workanator/vuego.v1/view"
)

type router struct{}

func (router) Find(sess *session.Session, route string) (v *view.View, err error) {
	content := &vuetify.App{
		Tag: ui.Tag{
			Id: "app",
		},
		Appearance: vuetify.Dark,
		Children: layout.Vert{
			&ui.Html{
				Content: "<p>Application says 'message'</p>",
				Events: ui.Listeners{
					ui.OnMouseMove: ui.Handle(func(sess *session.Session, cmp ui.Component, data interface{}) error {
						println("Clicked text")
						return nil
					}),
				},
			},
		},
	}

	return &view.View{
		Template: view.TemplatePage,
		Title:    "To-Do List",
		Content:  content,
	}, nil
}