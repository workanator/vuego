package model

type FieldModel struct {
	Owner PathModeler
	Path  []string
}

func (m *FieldModel) Model() interface{} {
	if m.Owner == nil {
		return nil
	}

	return m.Owner.PathModel(m.Path)
}

func (m *FieldModel) SetModel(value interface{}) {
	if m.Owner != nil {
		m.Owner.SetPathModel(m.Path, value)
	}
}
