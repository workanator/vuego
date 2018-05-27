package server

import (
	"golang.org/x/net/websocket"
)

func (server *Server) routeWs(conn *websocket.Conn) {
	// Decide how to interact with the client based on the protocol requested.
	protocol := conn.Config().Protocol
	if len(protocol) > 0 {
		// Connect client Write and server Read endpoints.
		if protocol[0] == "Model.Write" {
			server.wsModelRead(conn)
			return
		}

		// Connect client Read and server Write endpoints.
		if protocol[0] == "Model.Read" {
			server.wsModelWrite(conn)
			return
		}
	}

	// The protocol is not defined or unsupported
	conn.WriteClose(1002)
}
