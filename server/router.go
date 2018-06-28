package server

import (
	"io"
	"net/http"
	"strings"

	_ "gopkg.in/workanator/vuego.v1/resource"

	"github.com/sirupsen/logrus"
	"gopkg.in/workanator/vuego.v1/action"
	"gopkg.in/workanator/vuego.v1/session"
)

const (
	ActionDelimiter = ":"
	RouteApp        = "app"
	RouteAppAction  = RouteApp + ActionDelimiter
	RouteBus        = "bus"
	RouteStatic     = "static"
)

// Implement http.Handler interface.
func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	server.log.WithFields(logrus.Fields{
		"$M": r.Method,
		"$U": r.RequestURI,
	}).Debug("Request")

	// Resolve the session
	sess, err := server.bundle.Sessions.Resolve(r)
	if err != nil {
		if session.IsAccessDenied(err) {
			server.log.Error("Access Denied")
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("403 Hey! You shall not pass!"))
		} else {
			server.log.WithError(err).Error("Failed to resolve session")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 " + err.Error()))
		}

		return
	}

	// Sanitize the request URI and split it to segments
	uri := strings.TrimFunc(r.RequestURI, func(r rune) bool {
		return r == '/'
	})
	route, act := server.decomposeUri(uri)
	segments := strings.Split(route, "/")

	// Should not happen
	if len(segments) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Server the request
	switch segments[0] {
	case "":
		// Redirect to application start screen
		to := "/" + RouteApp
		server.log.WithField("to", to).Debug("Redirect")
		w.Header().Set("Location", to)
		w.WriteHeader(http.StatusMovedPermanently)
		return

	case RouteApp:
		// Render screen or process action
		switch r.Method {
		case "GET":
			// GET method is for rendering content
			if err := server.renderAppScreen(w, sess, route); err != nil {
				server.renderError(w, r, err)
			}

		default:
			// Other methods are actions
			if act, err := action.Parse(act); err == nil {
				if err := server.respondAppAction(w, sess, act); err != nil {
					server.renderError(w, r, err) // TODO: Return JSON error
				}
			} else {
				server.renderError(w, r, err)
			}
		}

		return

	case RouteBus:
		server.ws.ServeHTTP(w, r)
		return

	case RouteStatic:
		if len(segments) > 1 {
			path := strings.Join(segments[1:], "/")

			if f, err := server.fs.Open(path); err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(err.Error()))
			} else {
				switch segments[1] {
				case "js":
					w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
				case "css":
					w.Header().Set("Content-Type", "text/css; charset=utf-8")
				case "html":
					w.Header().Set("Content-Type", "text/html; charset=utf-8")
				}

				w.WriteHeader(http.StatusOK)
				io.Copy(w, f)
			}

			return
		}

	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found :("))
}

func (server *Server) renderError(w http.ResponseWriter, r *http.Request, err error) {
	// Write the response
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func (server *Server) decomposeUri(s string) (route, action string) {
	if len(s) == 0 {
		return "", ""
	}

	if pos := strings.Index(s, ActionDelimiter); pos >= 0 {
		return s[:pos-1], s[pos+1:]
	} else {
		return s, ""
	}
}
