package app

import (
	"net/http"

	"gopkg.in/workanator/vuego.v1/app/route"
	"gopkg.in/workanator/vuego.v1/app/session"
)

type Bundle struct {
	Id        string
	Name      string
	Version   string
	Lifecycle Lifecycle
	Fs        http.FileSystem
	Sessions  session.Resolver
	Routes    route.Router
}
