package server

import "golang.org/x/net/websocket"

// The function works on protocol Bus.Write which is one way delivery protocol where events
// are written by client to server.
// Client -> Server
func (server *Server) wsModelRead(conn *websocket.Conn, sess *Session) {
	server.log.Info("Accept Bus.Write connection")

	// Start an infinite loop for reading model updates on client's side.
	for {
		// Test if the connection should be aborted
		select {
		case <-sess.Context.Done():
			break
		default:
		}

		// Read the message
		var message string
		if err := websocket.Message.Receive(conn, &message); err != nil {
			server.log.
				WithError(err).
				Error("Bus read failed")
			break
		}

		server.log.
			WithField("message", message).
			Debug("Bus read")
	}

	// Close the connection
	conn.WriteClose(1000)
}
