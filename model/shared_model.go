package model

import (
	"encoding/json"
)

// SharedModel represents a map object which data is shared and valid through screen session.
// The model is equivalent to the following code snippet.
//
//    new Vue({
//      data: {/*MODEL*/}
//    })
type SharedModel struct {
	BasicModel
}

func (m *SharedModel) Field(path ...string) ModelInitialer {
	return ModelInitial{
		Modeler: &FieldModel{
			Owner: m,
			Path:  path,
		},
	}
}

func (m *SharedModel) Markup() string {
	m.RLock()
	defer m.RUnlock()

	if data, err := json.Marshal(m.BasicModel.container); err != nil {
		s, _ := json.Marshal(err.Error())
		return "new Vue({data:{MODEL_ERROR:" + string(s) + "}})"
	} else {
		return "new Vue({data:" + string(data) + "})"
	}
}
