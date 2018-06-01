package server

import (
	"net/http"
	"strings"

	"fmt"

	"crypto/md5"

	"gopkg.in/workanator/vuego.v1/app"
	"gopkg.in/workanator/vuego.v1/html"
)

func (server *Server) handleApp(w http.ResponseWriter, r *http.Request, sess *Session, tpl []byte) error {
	// Get the top screen
	var screen app.Screener
	if len(sess.Screens) > 0 {
		screen = sess.Screens[len(sess.Screens)-1]
	}

	// Render application parts
	var (
		name       string
		title      string
		headHtml   string
		bodyHtml   string
		modelsHtml string
	)

	if screen != nil {
		// Get name and title
		name = screen.Name()
		title = screen.Title()

		// Render headers form the screen and add a meta tag with screen name
		headHtml = `<meta name="vuego:screen:name" content="` + name + `"/>`

		if headEl := screen.Head(); headEl != nil {
			if markup, err := headEl.Markup(); err != nil {
				return err
			} else {
				headHtml += "\n" + markup
			}
		}

		// Render body
		if bodyCmp := screen.Body(); bodyCmp != nil {
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

		// Render models in a script tag. Each model is put in a map and has a unique identifier
		// made from the screen name and the position of the model in the slice.
		models := screen.Models()
		if len(models) > 0 {
			modelsHtml = "<script>let Model = {};"

			for i, m := range models {
				if markup, err := m.Markup(); err != nil {
					return err
				} else {
					key := fmt.Sprintf("%s.%04d", name, i)
					encodedKey := fmt.Sprintf("%x", md5.Sum([]byte(key)))

					modelsHtml += "\n"
					modelsHtml += "Model['" + encodedKey + "']=" + markup + ";"
				}
			}

			modelsHtml += "</script>"
		}
	}

	// Process the template
	body := string(tpl)
	body = strings.Replace(body, "#TITLE#", title, -1)
	body = strings.Replace(body, "#HEAD#", headHtml, -1)
	body = strings.Replace(body, "#BODY#", bodyHtml, -1)
	body = strings.Replace(body, "#MODEL#", modelsHtml, -1)

	// Write the response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))

	return nil
}
