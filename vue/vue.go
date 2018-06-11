package vue

import (
	"encoding/json"
	"strings"

	"gopkg.in/workanator/vuego.v1/html"
	"gopkg.in/workanator/vuego.v1/model"
)

type Vue struct {
	Id   html.Id
	Data model.ModelPropertier
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

func (v *Vue) Property(name string) interface{} {
	if v.Data != nil {
		return v.Data.Property(name)
	}

	return nil
}

func (v *Vue) SetProperty(name string, value interface{}) {
	if v.Data != nil {
		v.Data.SetProperty(name, value)
	}
}

func (v *Vue) Markup() (string, error) {
	// Build Vue instance initializer
	sb := strings.Builder{}
	props := 0

	// Open the initializer
	sb.WriteString("new Vue({")

	// Add el property
	if len(v.Id) > 0 {
		el, _ := json.Marshal("#" + v.Id)
		sb.WriteString("el:")
		sb.WriteString(string(el))
		props++
	}

	// Add data property
	if v.Data != nil {
		if data := v.Data.Model(); data != nil {
			if json, err := json.Marshal(data); err != nil {
				return "", html.ErrMarkupFailed{
					Tag:    "Vue",
					Reason: err,
				}
			} else {
				if props > 0 {
					sb.WriteRune(',')
				}

				sb.WriteString("data:")
				sb.WriteString(string(json))
			}
		}
	}

	// Close the initializer
	sb.WriteString("})")

	return sb.String(), nil
}
