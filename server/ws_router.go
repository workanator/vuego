package server

import (
	"golang.org/x/net/websocket"
	"gopkg.in/workanator/vuego.v1/session"
)

func (server *Server) routeWs(conn *websocket.Conn) {
	// Identify the session
	sess, err := server.identifySession(conn.Request())
	if err != nil {
		if session.IsAccessDenied(err) {
			server.log.Error("Access Denied")
			conn.WriteClose(1008)
		} else {
			server.log.WithError(err).Error("Failed to identify session")
			conn.WriteClose(1011)
		}

		return
	}

	// Decide how to interact with the client based on the protocol requested.
	protocol := conn.Config().Protocol
	if len(protocol) > 0 {
		// Connect client Write and server Read endpoints.
		if protocol[0] == "Bus.Write" {
			server.wsModelRead(conn, sess)
			return
		}

		// Connect client Read and server Write endpoints.
		if protocol[0] == "Bus.Read" {
			server.wsModelWrite(conn, sess)
			return
		}
	}

	// The protocol is not defined or unsupported
	conn.WriteClose(1002)
}
