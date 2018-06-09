package app

import (
	"net/http"

	"gopkg.in/workanator/vuego.v1/action"
	"gopkg.in/workanator/vuego.v1/session"
)

type Bundle struct {
	Id        string
	Name      string
	Version   string
	Lifecycle Lifecycle
	Fs        http.FileSystem
	Sessions  session.Resolver
	Actions   action.Responder
}
