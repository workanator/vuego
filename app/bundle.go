package app

import (
	"net/http"

	"gopkg.in/workanator/vuego.v1/app/session"
	"gopkg.in/workanator/vuego.v1/mvc/view"
)

type Bundle struct {
	Id        string
	Name      string
	Version   string
	Lifecycle Lifecycle
	Fs        http.FileSystem
	Sessions  session.Resolver
	Screens   view.Representationer
}
