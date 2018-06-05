package server

import (
	"encoding/json"

	"golang.org/x/net/websocket"
	"gopkg.in/workanator/vuego.v1/event"
	"gopkg.in/workanator/vuego.v1/session"
)

const (
	eventBufSize = 64
)

// The function works on protocol Bus.Read which is one way delivery protocol where events
// are written by server to client.
// Server -> Client
func (server *Server) wsEventWrite(conn *websocket.Conn, sess *session.Session) {
	// Close the connection if outbound event bus is nil
	if sess.Outbound == nil {
		server.log.
			WithField("error", "outbound bus is nil").
			Error("Failed  to accept Bus.Read connection")
		conn.WriteClose(WsInternalError)
		return
	}

	// Accept the connection and
	server.log.Info("Accept Bus.Read connection")

	// Start an infinite loop for writing model updates on server's side.
	buf := make([]event.Event, eventBufSize)
	for {
		if n, err := sess.Outbound.Produce(buf); err != nil {
			server.log.WithError(err).Error("Event pull failed")
			break
		} else if n > 0 {
			for i := 0; i < n; i++ {
				// Encode event to JSON and send it
				if payload, err := json.Marshal(buf[i]); err != nil {
					server.log.WithError(err).Error("Event enconding failed")
				} else if err = websocket.Message.Send(conn, string(payload)); err != nil {
					server.log.WithError(err).Error("Event delivery failed")
				} else {
					server.log.WithField("payload", string(payload)).Debug("Event sent")
				}
			}
		} else {
			// Sounds like the event bus is disconnected
			server.log.Debug("No events received")
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
	conn.WriteClose(WsNormalClosure)
}
