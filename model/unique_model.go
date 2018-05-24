package model

import (
	"encoding/json"
)

// UniqueModel represents a map object which data is shared and valid through screen session.
// The model is equivalent to the following code snippet.
//
//    new Vue({
//      data: function() { return {/*MODEL*/} }
//    })
type UniqueModel struct {
	BasicModel
}

func (m *UniqueModel) Field(path ...string) ModelInitialer {
	return ModelInitial{
		Modeler: &FieldModel{
			Owner: m,
			Path:  path,
		},
	}
}

func (m *UniqueModel) Markup() string {
	m.RLock()
	defer m.RUnlock()

	if data, err := json.Marshal(m.BasicModel.container); err != nil {
		s, _ := json.Marshal(err.Error())
		return "new Vue({data:{MODEL_ERROR:" + string(s) + "}})"
	} else {
		return "new Vue({data:function(){return " + string(data) + "}})"
	}
}
