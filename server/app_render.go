package server

import (
	"net/http"
	"strings"

	"gopkg.in/workanator/vuego.v1/html"
	"gopkg.in/workanator/vuego.v1/session"
)

func (server *Server) renderAppScreen(w http.ResponseWriter, sess *session.Session, route string, tpl []byte) error {
	// Get the screen
	screen, err := server.bundle.Routes.Screen(sess, route)
	if err != nil {
		return err
	}

	// Render application parts
	var (
		bodyHtml string
	)

	if screen != nil {
		// Render body
		if screen.Root != nil {
			if bodyEl, err := screen.Root.Render(nil, html.Rect{}); err != nil {
				return err
			} else if bodyEl != nil {
				if markup, err := bodyEl.Markup(); err != nil {
					return err
				} else {
					bodyHtml = markup
				}
			}
		}
	}

	// Process the template
	body := string(tpl)
	body = strings.Replace(body, "#SCREEN.TITLE#", screen.Title, -1)
	body = strings.Replace(body, "#SCREEN.ROOT#", bodyHtml, -1)

	// Write the response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))

	return nil
}
