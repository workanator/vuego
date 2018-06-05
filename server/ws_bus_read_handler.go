package server

import (
	"encoding/json"

	"golang.org/x/net/websocket"
	"gopkg.in/workanator/vuego.v1/app/event"
	"gopkg.in/workanator/vuego.v1/app/session"
)

// The function works on protocol Bus.Write which is one way delivery protocol where events
// are written by client to server.
// Client -> Server
func (server *Server) wsModelRead(conn *websocket.Conn, sess *session.Session) WsCloseCode {
	// Close the connection if inbound event bus is nil
	if sess.Inbound == nil {
		server.log.
			WithField("error", "inbound bus is nil").
			Error("Failed  to accept Bus.Write connection")
		return WsInternalError
	}

	// Accept the connection
	server.log.Info("Accept Bus.Write connection")

	// Start an infinite loop for reading model updates on client's side.
	for {
		// Read the message
		// TODO: Make Receive be session context aware. Possible?
		var msg string
		if err := websocket.Message.Receive(conn, &msg); err != nil {
			server.log.
				WithError(err).
				Error("Bus.Write read failed")
			return WsTryAgainLater
		}

		server.log.
			WithField("payload", msg).
			Debug("Bus.Write event read")

		// Unmarshal payload
		var ev event.Event
		if err := json.Unmarshal([]byte(msg), &ev); err != nil {
			server.log.
				WithError(err).
				Error("Bus.Write JSON decode failed")
			continue
		}

		// Push the event to the event bus
		if err := sess.Inbound.Consume([]event.Event{ev}, sess.Context); err != nil {
			server.log.
				WithError(err).
				Error("Bus.Write event consume failed")
			return WsTryAgainLater
		}

		// Test if the connection should be closed
		if sess.Context != nil {
			select {
			case <-sess.Context.Done():
				break
			default:
			}
		}
	}

	return WsNormalClosure
}
