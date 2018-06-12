package server

import (
	"net/http"

	"gopkg.in/workanator/vuego.v1/session"
)

func (server *Server) renderAppScreen(w http.ResponseWriter, sess *session.Session, route string) error {
	// Find the view to render
	view, err := server.bundle.Views.Find(sess, route)
	if err != nil {
		return err
	}

	// Render empty response if the view returned is nil
	if view == nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
	}

	// Render the view
	return view.Render(server.fs, WriteFunc(func(b []byte) (int, error) {
		// The template is successfully rendered so 200 OK status can be sent
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		return w.Write(b)
	}))
}
