package model

type FieldModel struct {
	Accessor Propertier
	Name     string
}

func (m *FieldModel) Model() interface{} {
	if m.Accessor != nil {
		return m.Accessor.Property(m.Name)
	}

	return nil
}

func (m *FieldModel) SetModel(value interface{}) {
	if m.Accessor != nil {
		m.Accessor.SetProperty(m.Name, value)
	}
}
