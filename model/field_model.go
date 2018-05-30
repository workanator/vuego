package model

type FieldModel struct {
	Container GetSeter
	Name      string
}

func (m *FieldModel) Model() interface{} {
	if m.Container != nil {
		return m.Container.Get(m.Name)
	}

	return nil
}

func (m *FieldModel) SetModel(value interface{}) {
	if m.Container != nil {
		m.Container.Set(m.Name, value)
	}
}
