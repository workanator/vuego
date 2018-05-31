package vue

import (
	"encoding/json"
	"strings"

	"gopkg.in/workanator/vuego.v1/errors"
	"gopkg.in/workanator/vuego.v1/model"
)

type Vue struct {
	Id   string
	Data model.ModelMarkuper
}

func (v *Vue) Model() interface{} {
	if v.Data != nil {
		return v.Data.Model()
	}

	return nil
}

func (v *Vue) SetModel(value interface{}) {
	if v.Data != nil {
		v.Data.SetModel(value)
	}
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
