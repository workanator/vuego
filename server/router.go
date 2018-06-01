package server

import (
	"io"
	"net/http"
	"strings"

	_ "gopkg.in/workanator/vuego.v1/resource"

	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/workanator/vuego.v1/session"
)

// Implement http.Handler interface.
func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	server.log.WithFields(logrus.Fields{
		"$M": r.Method,
		"$U": r.RequestURI,
	}).Debug("Request")

	// Identify the session
	sess, err := server.identifySession(r)
	if err != nil {
		if session.IsAccessDenied(err) {
			server.log.Error("Access Denied")
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("403 Hey! You shall not pass!"))
		} else {
			server.log.WithError(err).Error("Failed to identify session")
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

	if len(segments) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Server the request
	switch segments[0] {
	case "app":
		// Load application template and pass the handler
		if tpl, err := server.readFileContent("html/app.html"); err != nil {
			server.renderError(w, r, err)
		} else if err := server.handleApp(w, r, sess, tpl); err != nil {
			server.renderError(w, r, err)
		}

		return

	case "ws":
		server.ws.ServeHTTP(w, r)
		return

	case "static":
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

func (server *Server) identifySession(r *http.Request) (*Session, error) {
	// Identify the user
	var user *session.User
	if server.bundle.SessionIdentifier != nil {
		if u, err := server.bundle.SessionIdentifier.Identify(r); err != nil {
			return nil, err
		} else {
			user = u
		}
	}

	// Start new session if required
	var sess *Session
	if len(server.sessions) == 0 {
		if s, err := NewSession(user, server.bundle.StartScreen); err != nil {
			return nil, err
		} else {
			sess = s
		}

		if user != nil {
			server.log.WithField("user", *user).Debug("Identified user")
		}

		server.log.Debug("Started new session")
		server.sessions = append(server.sessions, sess)
	} else {
		sess = server.sessions[0] // TODO: replace with session search
	}

	return sess, nil
}

func (server *Server) renderError(w http.ResponseWriter, r *http.Request, err error) {
	// Write the response
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func (server *Server) readFileContent(path string) ([]byte, error) {
	// Open the application template
	f, err := server.fs.Open("html/app.html")
	if err != nil {
		return nil, err
	}

	// Read the whole content of the template
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return content, nil
}
