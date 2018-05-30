package model

import (
	"sync"
)

// Container adds thread safe guarantees over Container.
type Container struct {
	sync.RWMutex
	data map[string]interface{}
}

func (m *Container) Get(name string) interface{} {
	m.RLock()
	defer m.RUnlock()

	if m.data != nil {
		return m.data[name]
	}

	return nil
}

func (m *Container) Set(name string, value interface{}) {
	m.Lock()
	defer m.Unlock()

	if m.data == nil {
		m.data = make(map[string]interface{})
	}

	m.data[name] = value
}
