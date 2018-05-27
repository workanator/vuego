package server

import "golang.org/x/net/websocket"

func (server *Server) wsModelRead(conn *websocket.Conn) {
	server.log.Info("Accept Model Write connection")

	// Start an infinite loop for reading model updates on client's side.
	for {
		// Read the message
		var message string
		if err := websocket.Message.Receive(conn, &message); err != nil {
			server.log.
				WithError(err).
				Error("Model read failed")
			break
		}

		server.log.
			WithField("message", message).
			Debug("Read model")
	}
}
