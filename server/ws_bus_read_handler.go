package server

import (
	"golang.org/x/net/websocket"
	"gopkg.in/workanator/vuego.v1/event"
	"gopkg.in/workanator/vuego.v1/session"
)

// The function works on protocol Bus.Write which is one way delivery protocol where events
// are written by client to server.
// Client -> Server
func (server *Server) wsModelRead(conn *websocket.Conn, sess *session.Session) {
	// Close the connection if inbound event bus is nil
	if sess.Inbound == nil {
		server.log.
			WithField("error", "inbound bus is nil").
			Error("Failed  to accept Bus.Write connection")
		conn.WriteClose(WsInternalError)
		return
	}

	// Accept the connection
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
		var ev event.Event
		if err := websocket.Message.Receive(conn, &ev); err != nil {
			server.log.
				WithError(err).
				Error("Bus read failed")
			break
		}

		server.log.
			WithField("event", &ev).
			Debug("Bus read")

		// Push the event to the event bus
		sess.Inbound.Consume([]event.Event{ev})
	}

	// Close the connection
	conn.WriteClose(WsNormalClosure)
}
