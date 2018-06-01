package server

import (
	"time"

	"golang.org/x/net/websocket"
)

func (server *Server) wsModelWrite(conn *websocket.Conn, sess *Session) {
	server.log.Info("Accept Bus Read connection")

	// Start an infinite loop for writing model updates on server's side.
	for {
		// Test if the connection should be aborted
		select {
		case <-sess.Context.Done():
			break
		default:
		}

		// Write the message
		websocket.Message.Send(conn, time.Now().String())

		time.Sleep(250 * time.Millisecond)
	}

	// Close the connection
	conn.WriteClose(1000)
}
