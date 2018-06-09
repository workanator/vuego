package server

import (
	"io"
	"net/http"
	"strings"

	_ "gopkg.in/workanator/vuego.v1/resource"

	"github.com/sirupsen/logrus"
	"gopkg.in/workanator/vuego.v1/session"
)

const (
	RouteApp    = "app"
	RouteBus    = "bus"
	RouteStatic = "static"
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
	segments := strings.Split(uri, "/")

	// Redirect to /app
	if len(segments) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Server the request
	switch segments[0] {
	case "":
		to := "/" + RouteApp
		server.log.WithField("to", to).Debug("Redirect")
		w.Header().Set("Location", to)
		w.WriteHeader(http.StatusMovedPermanently)
		return

	case RouteApp:
		// Build application action
		action := ""
		if len(segments) > 1 {
			action = strings.Join(segments[1:], "/")
		}

		// Load application template and pass the handler
		if tpl, err := server.readFileContent("html/app.html"); err != nil {
			server.renderError(w, r, err)
		} else if err := server.handleApp(w, r, sess, action, tpl); err != nil {
			server.renderError(w, r, err)
		}
		return

	case RouteBus:
		server.ws.ServeHTTP(w, r)
		return

	case RouteStatic:
		if len(segments) > 1 {
			path := strings.Join(segments[1:], "/")

			if f, err := server.openFile(path); err != nil {
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
