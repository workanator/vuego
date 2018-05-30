package vue

import (
	"encoding/json"
	"strings"

	"gopkg.in/workanator/vuego.v1/errors"
	"gopkg.in/workanator/vuego.v1/html"
)

type Vue struct {
	Id   string
	Data html.Markuper
}

func (v *Vue) Markup() (string, error) {
	// Build Vue instance initializer
	sb := strings.Builder{}

	// Open the initializer
	sb.WriteString("new Vue({")

	// Add el property
	if len(v.Id) > 0 {
		el, _ := json.Marshal("#" + v.Id)
		sb.WriteString("el:")
		sb.WriteString(string(el))
		sb.WriteRune(',')
	}

	// Add data property
	if v.Data != nil {
		if markup, err := v.Data.Markup(); err != nil {
			return "", errors.ErrMarkupFailed{
				Tag:    "Vue",
				Reason: err,
			}
		} else {
			sb.WriteString("data:")
			sb.WriteString(markup)
		}
	}

	// Close the initializer
	sb.WriteString("})")

	return sb.String(), nil
}
