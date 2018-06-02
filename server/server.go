package server

import (
	"net/http"

	"net"

	"fmt"

	"github.com/phogolabs/parcello"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
	"gopkg.in/workanator/vuego.v1/app"
	_ "gopkg.in/workanator/vuego.v1/resource"
)

type Server struct {
	ListenIP   net.IP
	ListenPort uint16
	bundle     app.Bundle
	fs         http.FileSystem
	log        *logrus.Entry
	ws         *websocket.Server
}

// Start prepares the server instance and starts listen for incoming requests. If some fields in the struct omitted
// they are initialized with default values. In most cases it's enough to provide valid ListenIP and ListenPort
// to start the server. When server started it blocks further execution of the current goroutine.
func (server Server) Start(bundle app.Bundle) error {
	server.bundle = bundle

	// Prepare the instance for start.
	if err := server.prepare(); err != nil {
		return err
	}

	// Start the server
	listenAddr := fmt.Sprintf("%s:%d", server.ListenIP.String(), server.ListenPort)
	server.log.WithField("listen_addr", listenAddr).Info("Starting server")
	err := http.ListenAndServe(listenAddr, &server)

	// Ignore server closed error.
	if err != http.ErrServerClosed {
		return err
	}

	return nil
}

// Prepares the server for start.
func (server *Server) prepare() error {
	// Initialize private members
	server.fs = parcello.Root("/")
	server.log = logrus.NewEntry(logrus.StandardLogger())

	// Initialize WebSocket server.
	server.ws = &websocket.Server{
		Handler: server.routeWs,
	}

	return nil
}
