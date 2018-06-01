package server

import (
	"encoding/json"

	"golang.org/x/net/websocket"
	"gopkg.in/workanator/vuego.v1/event"
)

// The function works on protocol Bus.Read which is one way delivery protocol where events
// are written by server to client.
// Server -> Client
func (server *Server) wsEventWrite(conn *websocket.Conn, sess *Session) {
	var cons event.PullQueue

	// Accept connection only if the bus in disconnected state
	accepted := false
	if sess.OutboundEvents.Bus.IsDisconnected() {
		cons = event.NewPullQueue(server.OutboundQueueSize)
		defer cons.Close()

		if err := sess.OutboundEvents.Bus.Connect(sess.OutboundEvents.Pusher, &cons); err != nil {
			server.log.WithError(err).Error("Failed to connect outbound events bus")
			conn.WriteClose(1011)
		} else {
			accepted = true
		}
	} else {
		server.log.Error("Outbound events bus is already connected")
		conn.WriteClose(1008)
	}

	// Continue only the connection is accepted
	if accepted {
		server.log.Info("Bus.Read connection is accepted")
	} else {
		server.log.Info("Bus.Read connection is rejected")
		return
	}

	// Disconnect the bus when finished
	defer sess.OutboundEvents.Bus.Disconnect()

	// Start an infinite loop for writing model updates on server's side.
	for {
		// Test if the connection should be aborted
		select {
		case e := <-cons.Queue():
			if e != nil {
				// Encode event to JSON and send it
				if payload, err := json.Marshal(*e); err != nil {
					server.log.WithError(err).Error("Event econding failed")
				} else if err = websocket.Message.Send(conn, payload); err != nil {
					server.log.WithError(err).Error("Event delivery failed")
				} else {
					server.log.WithField("payload", string(payload)).Debug("Event sent")
				}
			} else {
				// Sounds like the queue is closed
				server.log.Debug("Nil event received")
				break
			}
		case <-sess.Context.Done():
			break
		default:
		}
	}

	// Close the connection
	conn.WriteClose(1000)
}
