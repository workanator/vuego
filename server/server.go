package server

import (
	"net/http"

	"github.com/sirupsen/logrus"
	_ "gopkg.in/workanator/vuego.v1/resource"
	"gopkg.in/workanator/vuego.v1/ui"
)

type Server struct {
	Router  http.Handler
	Logger  *logrus.Entry
	Screens []ui.Screener
}

// DefaultServer creates Server instance with default values.
func DefaultServer() *Server {
	return &Server{
		Router:  DefaultRouter(),
		Logger:  logrus.NewEntry(logrus.StandardLogger()),
		Screens: []ui.Screener{},
	}
}

func (server *Server) Start(screen ui.Screener) error {
	server.Logger.Info("Starting server")

	return http.ListenAndServe(":8008", server.Router)
}
