package test_todo

import (
	"net/http"

	"strings"

	"context"

	"gopkg.in/workanator/vuego.v1/app"
	"gopkg.in/workanator/vuego.v1/app/facade"
	"gopkg.in/workanator/vuego.v1/app/session"
)

func Bundle() app.Bundle {
	reactor := facade.NewReactor()

	return app.Bundle{
		Id:      "todo",
		Name:    "Simple To-Do",
		Version: "1.0",
		Fs:      &facade.MultiFS{},
		Sessions: &facade.MultiSession{
			Who: session.IdentifyFunc(func(r *http.Request) (string, error) {
				return (strings.Split(r.RemoteAddr, ":"))[0], nil
			}),
			Store: &facade.SessionPool{
				New: session.StartFunc(func(sessId string) (*session.Session, error) {
					return &session.Session{
						Context: context.Background(),
						Id:      sessId,
						User:    nil,
						State:   nil,
					}, nil
				}),
			},
			Inbound:  reactor,
			Outbound: reactor,
		},
		Screens: &reprManager{},
	}
}
