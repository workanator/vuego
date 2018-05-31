package model

import (
	"sync"
)

type dataContainer = map[string]interface{}

// Container implements the basic thread-safe model. It implements Modeler, Propertier and Markuper interfaces
// so it can be used with Screener interface.
type Container struct {
	sync.RWMutex
	data dataContainer
}

// Construct FieldModel instance. The type returned is ModelInitialer so the model can be configured and initialized
// with value.
func (m *Container) Field(name string) ModelInitialer {
	defer m.SetProperty(name, nil)

	return ModelInitial{
		Modeler: &FieldModel{
			Accessor: m,
			Name:     name,
		},
	}
}

func (m *Container) Model() interface{} {
	m.Lock()
	defer m.Unlock()

	if m.data == nil {
		m.data = make(dataContainer)
	}

	return m.data
}

func (m *Container) SetModel(value interface{}) {
	m.Lock()
	defer m.Unlock()

	if value != nil {
		// Copy values from the value
		switch c := value.(type) {
		case dataContainer:
			if m.data == nil {
				m.data = make(dataContainer)
			}

			for key := range c {
				m.data[key] = c[key]
			}

		case *dataContainer:
			if m.data == nil {
				m.data = make(dataContainer)
			}

			for key := range *c {
				m.data[key] = (*c)[key]
			}
		}
	} else {
		// Erase container data
		m.data = make(dataContainer)
	}
}

func (m *Container) Property(name string) interface{} {
	m.RLock()
	defer m.RUnlock()

	if m.data != nil {
		return m.data[name]
	}

	return nil
}

func (m *Container) SetProperty(name string, value interface{}) {
	m.Lock()
	defer m.Unlock()

	if m.data == nil {
		m.data = make(dataContainer)
	}

	m.data[name] = value
}
