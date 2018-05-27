package server

import (
	"time"

	"golang.org/x/net/websocket"
)

func (server *Server) wsModelWrite(conn *websocket.Conn) {
	server.log.Info("Accept Model Read connection")

	// Start an infinite loop for writing model updates on server's side.
	for {
		// Write the message
		websocket.Message.Send(conn, time.Now().String())

		time.Sleep(250 * time.Millisecond)
	}
}
