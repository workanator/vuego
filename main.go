package main

import (
	"net"
	"sync/atomic"

	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
	"gopkg.in/workanator/vuego.v1/browser"
	"gopkg.in/workanator/vuego.v1/server"
	"gopkg.in/workanator/vuego.v1/test_todo"
)

func main() {
	// Configure the logger
	logrus.SetLevel(logrus.DebugLevel)

	// Track lifetime of the parts of the application.
	var (
		partsRunning     int32
		serverErrorChan  = make(chan error)
		browserErrorChan = make(chan error)
	)

	defer close(browserErrorChan)
	defer close(serverErrorChan)

	// Start the server
	var serv atomic.Value
	atomic.AddInt32(&partsRunning, 1)
	go func() {
		defer atomic.AddInt32(&partsRunning, -1)

		srv := &server.Server{
			ListenIP:   net.ParseIP("127.0.0.1"),
			ListenPort: 8008,
		}
		serv.Store(srv)

		err := srv.Start(test_todo.Bundle())

		// Send the error.
		serverErrorChan <- err
	}()

	// Start the browser
	atomic.AddInt32(&partsRunning, 1)
	go func() {
		defer atomic.AddInt32(&partsRunning, -1)

		err := browser.Lauch(
			"http://127.0.0.1:8008/app",
			&browser.Options{
				NewInstance: false,
				Window: browser.WindowOptions{
					Width: 800,
				},
			},
			browser.Firefox(),
		)

		// Send the error.
		browserErrorChan <- err
	}()

	// React on OS signals.
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

	// Wait until one part is finished.
	interrupt := false
	for !interrupt {
		select {
		case err := <-browserErrorChan:
			if err != nil {
				logrus.WithError(err).Error("Browser error")
			} else {
				logrus.Info("Browser finished")
			}

		case err := <-serverErrorChan:
			if err != nil {
				logrus.WithError(err).Error("Server error")
			} else {
				logrus.Info("Server finished")
			}

		case signal := <-sigint:
			logrus.WithField("signal", signal).Info("Caught signal")
			interrupt = true

			if srv, ok := serv.Load().(*server.Server); ok && srv != nil {
				srv.Stop()
			}
		}

		// Check if all parts finished
		if atomic.LoadInt32(&partsRunning) == 0 {
			interrupt = true
		}
	}
}
