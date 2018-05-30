package model

import (
	"encoding/json"

	"gopkg.in/workanator/vuego.v1/errors"
)

// UniqueData represents a map object which data is shared and valid through screen session.
// The model is equivalent to the following code snippet.
//
//    new Vue({
//      data: function() { return {/*MODEL*/} }
//    })
type UniqueData Container

func (m *UniqueData) Field(name string) ModelInitialer {
	return ModelInitial{
		Modeler: &FieldModel{
			Container: (*Container)(m),
			Name:      name,
		},
	}
}

func (m *UniqueData) Markup() (string, error) {
	m.RLock()
	defer m.RUnlock()

	if data, err := json.Marshal(m.data); err != nil {
		return "", errors.ErrMarkupFailed{
			Tag:    "script",
			Reason: err,
		}
	} else {
		return "function(){return " + string(data) + "}", nil
	}
}
