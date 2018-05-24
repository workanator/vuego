package model

type FieldModel struct {
	Container PathModeler
	Path      []string
}

func (m *FieldModel) Model() interface{} {
	return m.Container.PathModel(m.Path)
}

func (m *FieldModel) SetModel(value interface{}) {
	m.Container.SetPathModel(m.Path, value)
}
