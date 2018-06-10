package test_todo

import (
	"net/http"

	"strings"

	"context"

	"gopkg.in/workanator/vuego.v1/app"
	"gopkg.in/workanator/vuego.v1/app/facade"
	"gopkg.in/workanator/vuego.v1/session"
)

func Bundle() app.Bundle {
	reactor := facade.NewReactor()

	return app.Bundle{
		Id:      "todo",
		Name:    "Simple To-Do",
		Version: "1.0",
		Fs:      &facade.MultiFS{},
		Lifecycle: app.LifecycleFunc(func(starting bool, b *app.Bundle) error {
			return nil
		}),
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
						Data:    nil,
					}, nil
				}),
			},
			EventBus: reactor,
		},
		Routes: &router{},
	}
}
