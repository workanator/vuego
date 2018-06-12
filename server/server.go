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
	fs         multiFs
	log        *logrus.Entry
	serv       *http.Server
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

	// Create and configure HTTP server instance
	listenAddr := fmt.Sprintf("%s:%d", server.ListenIP.String(), server.ListenPort)
	server.serv = &http.Server{
		Addr:    listenAddr,
		Handler: &server,
	}

	// Startup application
	if bundle.Lifecycle != nil {
		if err := bundle.Lifecycle.Startup(&bundle); err != nil {
			return err
		}
	}

	server.serv.RegisterOnShutdown(func() {
		// Shutdown application
		if bundle.Lifecycle != nil {
			if err := bundle.Lifecycle.Shutdown(&bundle); err != nil {
				server.log.WithError(err).Error("")
			}
		}

		server.log.WithField("listen_addr", listenAddr).Info("Stopping server")
	})

	// Start the server
	server.log.WithField("listen_addr", listenAddr).Info("Starting server")
	if err := server.serv.ListenAndServe(); err != nil {
		// Ignore server closed error.
		if err != http.ErrServerClosed {
			return err
		}
	}

	return nil
}

// Stop the running server.
func (server *Server) Stop() {
	if server.serv != nil {
		server.serv.Shutdown(nil)
	}
}

// Prepares the server for start.
func (server *Server) prepare() error {
	// Initialize private members
	server.fs = multiFs{parcello.Root("/")}
	if server.bundle.Fs != nil {
		server.fs = append(server.fs, server.bundle.Fs)
	}

	server.log = logrus.NewEntry(logrus.StandardLogger())

	// Initialize WebSocket server.
	server.ws = &websocket.Server{
		Handler: server.routeWs,
	}

	return nil
}
