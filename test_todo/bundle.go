package test_todo

import "gopkg.in/workanator/vuego.v1/app"

func Bundle() app.Bundle {
	return app.Bundle{
		Id:                    "todo",
		Name:                  "Simple To-Do",
		Version:               "1.0",
		Fs:                    nil,
		SessionManager:        &sessManager{},
		RepresentationManager: &reprManager{},
	}
}
