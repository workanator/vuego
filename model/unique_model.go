package model

import (
	"encoding/json"

	"gopkg.in/workanator/vuego.v1/errors"
)

// UniqueModel represents a map object which data is shared and valid through screen session.
// The model is equivalent to the following code snippet.
//
//    new Vue({
//      data: function() { return {/*MODEL*/} }
//    })
type UniqueModel struct {
	BasicModel
	Id string
}

func (m *UniqueModel) Field(path ...string) ModelInitialer {
	return ModelInitial{
		Modeler: &FieldModel{
			Owner: m,
			Path:  path,
		},
	}
}

func (m *UniqueModel) Markup() (string, error) {
	m.RLock()
	defer m.RUnlock()

	if data, err := json.Marshal(m.BasicModel.container); err != nil {
		return "", errors.ErrMarkupFailed{
			Tag:    "script",
			Reason: err,
		}
	} else {
		if len(m.Id) > 0 {
			id, _ := json.Marshal("#" + m.Id)
			return "new Vue({el:" + string(id) + ",data:function(){return " + string(data) + "}})", nil
		} else {
			return "new Vue({data:function(){return " + string(data) + "}})", nil
		}
	}
}