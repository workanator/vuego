package main

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/workanator/vuego.v1/browser"
	"gopkg.in/workanator/vuego.v1/server"
	"gopkg.in/workanator/vuego.v1/ui"
)

func main() {
	// Configure the logger
	logrus.SetLevel(logrus.DebugLevel)

	// Start the server
	go server.DefaultServer().Start(nil)

	// Start the browser
	browser.Lauch(
		"http://127.0.0.1:8008/app.html",
		&browser.LaunchOptions{
			Priority:    []browser.Launcher{browser.Firefox()},
			NewInstance: false,
			Window: &ui.WindowOptions{
				Size: &ui.BoxSize{
					Width: 800,
				},
			},
		},
	)
}
