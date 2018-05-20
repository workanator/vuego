package server

import (
	"io"
	"net/http"
	"strings"

	_ "gopkg.in/workanator/vuego.v1/resource"

	"io/ioutil"

	"github.com/phogolabs/parcello"
	"github.com/sirupsen/logrus"
	"gopkg.in/workanator/vuego.v1/html"
	"gopkg.in/workanator/vuego.v1/theme/vuetify"
	"gopkg.in/workanator/vuego.v1/ui"
)

type Router struct {
	StaticFS http.FileSystem
	Renderer html.Renderer
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
	case "app":
		f, err := router.StaticFS.Open("html/app.html")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
		}

		content, err := ioutil.ReadAll(f)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
		}

		body := string(content)
		e := &vuetify.App{
			Appearance: vuetify.Dark,
			Toolbar: struct {
				Top    *vuetify.Toolbar
				Bottom *vuetify.Toolbar
			}{
				Bottom: &vuetify.Toolbar{},
			},
			Children: ui.VerticalLayout{
				&ui.Text{
					Tag: ui.Tag{
						Style: html.Style{
							"border": "4px double black",
						},
					},
					Text: "{{ message }}",
					Type: ui.TextParagraph,
				},
				&ui.Text{
					Tag: ui.Tag{
						Style: html.Style{
							"border": "4px double black",
						},
					},
					Text: "{{ message }}",
					Type: ui.TextParagraph,
				},
			},
		}

		body = strings.Replace(body, "#BODY.BEFORE#", e.Render(nil, html.Rect{}).Markup(), -1)
		body = strings.Replace(body, "#BODY.AFTER#", "<script>var app = new Vue({el: '#app', data: {message: 'Zdarov, Vue!'}})</script>", -1)

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))

		return

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
