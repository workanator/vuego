package server

import (
	"io"
	"net/http"
	"strings"

	_ "gopkg.in/workanator/vuego.v1/resource"

	"io/ioutil"

	"github.com/phogolabs/parcello"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

type Router struct {
	StaticFS http.FileSystem
	log      *logrus.Entry
	wsServer *websocket.Server
	appFunc  func(http.ResponseWriter, *http.Request, []byte) error
	wsFunc   func(conn *websocket.Conn)
}

// DefaultRouter creates Router instance and initialize it with default values.
func DefaultRouter() *Router {
	// Create router instance
	router := &Router{
		StaticFS: parcello.Root("/"),
		log:      logrus.NewEntry(logrus.StandardLogger()),
	}

	// Create and configure WebSocket server
	//	wsConfig, err := websocket.NewConfig("127.0.0.1", "")
	//	if err != nil {
	//		panic(err)
	//	}

	router.wsServer = &websocket.Server{
		Handler: func(conn *websocket.Conn) {
			if router.wsFunc != nil {
				router.wsFunc(conn)
			} else {
				conn.WriteClose(1011) // 1014 stands for Bad Gateway
			}
		},
	}

	return router
}

// Implement HTTP handler interface to make the reouter able to serve HTTP requests.
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	router.log.WithFields(logrus.Fields{
		"$M": r.Method,
		"$U": r.RequestURI,
	}).Debug("Request")

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
		if router.appFunc != nil {
			// Load application template and pass the handler
			if tpl, err := router.readFileContent("html/app.html"); err != nil {
				router.renderError(w, r, err)
			} else if err := router.appFunc(w, r, tpl); err != nil {
				router.renderError(w, r, err)
			}

			return
		}

	case "ws":
		if router.wsServer != nil {
			router.wsServer.ServeHTTP(w, r)
			return
		}

	case "static":
		if len(segments) > 1 {
			path := strings.Join(segments[1:], "/")

			if f, err := router.StaticFS.Open(path); err != nil {
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

func (router *Router) renderError(w http.ResponseWriter, r *http.Request, err error) {
	// Write the response
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func (router *Router) readFileContent(path string) ([]byte, error) {
	// Open the application template
	f, err := router.StaticFS.Open("html/app.html")
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
