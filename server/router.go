package server

import (
	"io"
	"net/http"
	"strings"

	_ "gopkg.in/workanator/vuego.v1/resource"

	"io/ioutil"

	"github.com/phogolabs/parcello"
	"github.com/sirupsen/logrus"
	"gopkg.in/workanator/vuego.v1/app"
	"gopkg.in/workanator/vuego.v1/html"
)

type Router struct {
	StaticFS http.FileSystem
	Screen   app.Screener
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
		if err := router.renderApp(w, r); err != nil {
			router.renderError(w, r, err)
		}
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

func (router *Router) renderApp(w http.ResponseWriter, r *http.Request) error {
	// Load file content
	body, err := router.readFileContent("html/app.html")
	if err != nil {
		return err
	}

	// Render application parts
	var (
		title      string
		headHtml   string
		bodyHtml   string
		modelsHtml string
	)

	if router.Screen != nil {
		// Get title
		title = router.Screen.Title()

		// Render headers
		if headEl := router.Screen.Head(); headEl != nil {
			if markup, err := headEl.Markup(); err != nil {
				return err
			} else {
				headHtml = markup
			}
		}

		// Render body
		if bodyCmp := router.Screen.Body(); bodyCmp != nil {
			if bodyEl, err := bodyCmp.Render(nil, html.Rect{}); err != nil {
				return err
			} else if bodyEl != nil {
				if markup, err := bodyEl.Markup(); err != nil {
					return err
				} else {
					bodyHtml = markup
				}
			}
		}

		// Render modele
		for _, m := range router.Screen.Models() {
			if markup, err := m.Markup(); err != nil {
				return err
			} else {
				if len(modelsHtml) > 0 {
					modelsHtml += "\n"
				}

				modelsHtml += markup
			}
		}
	}

	// Process the template
	body = strings.Replace(body, "#TITLE#", title, -1)
	body = strings.Replace(body, "#HEAD#", headHtml, -1)
	body = strings.Replace(body, "#BODY#", bodyHtml, -1)
	body = strings.Replace(body, "#MODEL#", modelsHtml, -1)

	// Write the response
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))

	return nil
}

func (router *Router) renderError(w http.ResponseWriter, r *http.Request, err error) {
	// Write the response
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func (router *Router) readFileContent(path string) (string, error) {
	// Open the application template
	f, err := router.StaticFS.Open("html/app.html")
	if err != nil {
		return "", err
	}

	// Read the whole content of the template
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
