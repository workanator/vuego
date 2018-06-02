package server

import (
	"crypto/md5"
	"fmt"
	binHtml "html"
	"net/http"
	"strings"

	"gopkg.in/workanator/vuego.v1/html"
	"gopkg.in/workanator/vuego.v1/session"
)

func (server *Server) handleApp(w http.ResponseWriter, r *http.Request, sess *session.Session, tpl []byte) error {
	// Get the screen
	screen, err := server.bundle.RepresentationManager.Representation(sess)
	if err != nil {
		return err
	}

	// Render application parts
	var (
		headHtml   string
		bodyHtml   string
		modelsHtml string
	)

	if screen != nil {
		// Render headers form the screen and add a meta tag with screen name
		if len(screen.Id) > 0 {
			headHtml += `<meta name="vuego:screen:id" content="` + binHtml.EscapeString(screen.Id) + `"/>`
		}

		if len(screen.Name) > 0 {
			headHtml = `<meta name="vuego:screen:name" content="` + binHtml.EscapeString(screen.Name) + `"/>`
		}

		if screen.Head != nil {
			if markup, err := screen.Head.Markup(); err != nil {
				return err
			} else {
				headHtml += "\n" + markup
			}
		}

		// Render body
		if screen.Body != nil {
			if bodyEl, err := screen.Body.Render(nil, html.Rect{}); err != nil {
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
		if len(screen.Models) > 0 {
			modelsHtml = "<script>let Model = {};"

			for i, m := range screen.Models {
				if markup, err := m.Markup(); err != nil {
					return err
				} else {
					key := fmt.Sprintf("%s.%04d", screen.Name, i)
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
	body = strings.Replace(body, "#TITLE#", screen.Title, -1)
	body = strings.Replace(body, "#HEAD#", headHtml, -1)
	body = strings.Replace(body, "#BODY#", bodyHtml, -1)
	body = strings.Replace(body, "#MODEL#", modelsHtml, -1)

	// Write the response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))

	return nil
}
