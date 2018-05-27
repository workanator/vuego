package server

import (
	"net/http"

	"net"

	"fmt"

	"github.com/sirupsen/logrus"
	"gopkg.in/workanator/vuego.v1/app"
	_ "gopkg.in/workanator/vuego.v1/resource"
)

type Server struct {
	ListenIP   net.IP
	ListenPort uint16
	Router     *Router
	log        *logrus.Entry
	screens    []app.Screener
}

// Start prepares the server instance and starts listen for incoming requests. If some fields in the struct omitted
// they are initialized with default values. In most cases it's enough to provide valid ListenIP and ListenPort
// to start the server. When server started it blocks further execution of the current goroutine.
func (server Server) Start(screen app.Screener) error {
	// Prepare the instance for start.
	if err := server.prepare(); err != nil {
		return err
	}

	// Set the start screen
	server.PushScreen(screen)

	// Start the server
	listenAddr := fmt.Sprintf("%s:%d", server.ListenIP.String(), server.ListenPort)
	server.log.WithField("listen_addr", listenAddr).Info("Starting server")
	err := http.ListenAndServe(listenAddr, server.Router)

	// Ignore server closed error.
	if err != http.ErrServerClosed {
		return err
	}

	return nil
}

// Push the screen on top of the
func (server *Server) PushScreen(screen app.Screener) Refresher {
	if screen != nil {
		server.screens = append(server.screens, screen)
	}

	return nil
}

// Prepares the server for start.
func (server *Server) prepare() error {
	// Use teh default router if no router is provided.
	if server.Router == nil {
		server.Router = DefaultRouter()
		server.Router.appFunc = server.handleApp
		server.Router.wsFunc = server.handleWsConn
	}

	// Use the standard logger if no logger is provided.
	if server.log == nil {
		server.log = logrus.NewEntry(logrus.StandardLogger())
	}

	// Initialize screen stack.
	if server.screens == nil {
		server.screens = make([]app.Screener, 0)
	}

	return nil
}
