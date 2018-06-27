package server

import (
	"net/http"

	"encoding/json"

	"gopkg.in/workanator/vuego.v1/action"
	"gopkg.in/workanator/vuego.v1/session"
)

func (server *Server) respondAppAction(w http.ResponseWriter, sess *session.Session, act *action.Action) (err error) {
	// Respond action
	var response interface{}
	if server.bundle.Actions != nil {
		if response, err = server.bundle.Actions.Respond(sess, act); err != nil {
			return err
		}
	}

	// Marshal response to JSON
	var jsonData []byte
	if jsonData, err = json.Marshal(response); err != nil {
		return err
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

	return nil
}
