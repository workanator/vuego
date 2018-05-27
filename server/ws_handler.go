package server

import "golang.org/x/net/websocket"

func (server *Server) handleWsConn(conn *websocket.Conn) {
	server.log.Info("reading data")
	var s string
	websocket.Message.Receive(conn, &s)
	server.log.Info("received " + s)

	server.log.Info("writing data")
	websocket.Message.Send(conn, "HELLO")

	server.log.Info("closing")
	conn.Close()

	server.log.Info("closed")
}
