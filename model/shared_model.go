package model

import (
	"encoding/json"
)

// BasicModel represents a map object which data is shared and valid through screen session.
// The model is equivalent to
//
//    new Vue({
//      data: {}, // shared model
//    })
type SharedModel struct {
	BasicModel
}

func (m *SharedModel) Field(path ...string) ModelIniter {
	return modelInit{
		Modeler: &FieldModel{
			Container: m,
			Path:      path,
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

type modelInit struct {
	Modeler
}

func (mi modelInit) Init(value interface{}) Modeler {
	mi.Modeler.SetModel(value)
	return mi.Modeler
}
