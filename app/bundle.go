package app

import (
	"net/http"

	"gopkg.in/workanator/vuego.v1/session"
)

type Bundle struct {
	Id                    string
	Name                  string
	Version               string
	Fs                    http.FileSystem
	SessionManager        session.Resolver
	RepresentationManager Representationer
}
