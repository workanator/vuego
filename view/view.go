package view

import (
	"io"

	"net/http"

	"strings"

	"gopkg.in/workanator/vuego.v1/html"
	"gopkg.in/workanator/vuego.v1/ui"
)

type View struct {
	Template Path
	Title    string
	Content  ui.Component
}

// Load template from file system fs and render view into writer w.
func (v *View) Render(fs http.FileSystem, w io.Writer) (err error) {
	// Load template
	var tpl []byte
	if !v.Template.IsEmpty() {
		if tpl, err = v.Template.Load(fs); err != nil {
			return err
		} else if tpl == nil {
			return ErrNotFound{}
		}
	}

	// Render view content
	var contentMarkup string
	if v.Content != nil {
		if contentEl, err := v.Content.Render(nil, html.Rect{}); err != nil {
			return err
		} else if contentEl != nil {
			if contentMarkup, err = contentEl.Markup(); err != nil {
				return err
			}
		}
	}

	// Process the template and write to the writer
	body := string(tpl)
	if len(body) > 0 {
		body = strings.Replace(body, "#VIEW.TITLE#", v.Title, -1)
		body = strings.Replace(body, "#VIEW.CONTENT#", contentMarkup, -1)
	} else {
		body = contentMarkup
	}

	if n, err := io.WriteString(w, body); err != nil {
		return err
	} else if n != len(body) {
		return io.ErrShortWrite
	}

	return nil
}
