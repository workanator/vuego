package server

import (
	"io"
	"net/http"
	"strings"

	_ "gopkg.in/workanator/vuego.v1/resource"

	"github.com/phogolabs/parcello"
	"github.com/sirupsen/logrus"
)

type Router struct {
	StaticFS http.FileSystem
	log      *logrus.Entry
}

// DefaultRouter creates Router instance and initialize it with default values.
func DefaultRouter() *Router {
	return &Router{
		StaticFS: parcello.Root("/"),
		log:      logrus.NewEntry(logrus.StandardLogger()),
	}
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
	case "app.html":
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<!DOCTYPE html><html><head><script src=\"/static/js/vue.min.js\"></script></head><body>"))
		w.Write([]byte("<div id=\"app\">{{ message }}</div>"))
		w.Write([]byte("<script>var app = new Vue({el: '#app',data: {message: 'Zdarov, Vue!'}})</script>"))
		w.Write([]byte("</body></html>"))
		return

	case "static":
		if len(segments) > 1 {
			path := strings.Join(segments[1:], "/")

			if f, err := router.StaticFS.Open(path); err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(err.Error()))
			} else {
				w.Header().Set("Content-Type", "application/javascript")
				w.WriteHeader(http.StatusOK)
				io.Copy(w, f)
			}

			return
		}

	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found :("))
}
