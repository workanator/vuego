package server

import (
	"encoding/json"

	"golang.org/x/net/websocket"
	"gopkg.in/workanator/vuego.v1/session"
)

// The function works on protocol Bus.Read which is one way delivery protocol where events
// are written by server to client.
// Server -> Client
func (server *Server) wsEventWrite(conn *websocket.Conn, sess *session.Session) {
	// Start an infinite loop for writing model updates on server's side.
	for {
		if ev, err := sess.OutboundEvents.Pull(); err != nil {
			server.log.WithError(err).Error("Event pull failed")
			break
		} else if ev != nil {
			// Encode event to JSON and send it
			if payload, err := json.Marshal(ev); err != nil {
				server.log.WithError(err).Error("Event enconding failed")
			} else if err = websocket.Message.Send(conn, payload); err != nil {
				server.log.WithError(err).Error("Event delivery failed")
			} else {
				server.log.WithField("payload", string(payload)).Debug("Event sent")
			}
		} else {
			// Sounds like the event bus is disconnected
			server.log.Debug("Nil event received")
			break
		}

		// Test if the connection should be closed
		select {
		case <-sess.Context.Done():
			break
		default:
		}
	}

	// Close the connection
	conn.WriteClose(1000)
}
