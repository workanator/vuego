package server

import (
	"golang.org/x/net/websocket"
	"gopkg.in/workanator/vuego.v1/session"
)

func (server *Server) routeWs(conn *websocket.Conn) {
	// Resolve the session
	sess, err := server.bundle.Sessions.Resolve(conn.Request())
	if err != nil {
		if session.IsAccessDenied(err) {
			server.log.Error("Access Denied")
			conn.WriteClose(WsPolicyViolation)
		} else {
			server.log.WithError(err).Error("Failed to identify session")
			conn.WriteClose(WsInternalError)
		}

		return
	}

	// Decide how to interact with the client based on the protocol requested.
	protocol := conn.Config().Protocol
	if len(protocol) > 0 {
		server.log.WithField("protocol", protocol[0]).Debug("Bus connection request")

		// Connect client Write and server Read endpoints.
		if protocol[0] == "Bus.Write" {
			server.wsModelRead(conn, sess)
			return
		}

		// Connect client Read and server Write endpoints.
		if protocol[0] == "Bus.Read" {
			server.wsEventWrite(conn, sess)
			return
		}
	}

	// The protocol is not defined or unsupported
	conn.WriteClose(WsProtocolError)
}
