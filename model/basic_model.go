package model

import (
	"sync"
)

type mapContainer = map[string]interface{}

// BasicModel implements Modeler and PathModeler over go map structure.
type BasicModel struct {
	sync.RWMutex
	container mapContainer
}

func (m *BasicModel) Model() interface{} {
	m.RLock()
	defer m.RUnlock()

	return m.container
}

func (m *BasicModel) SetModel(value interface{}) {
	m.Lock()
	defer m.Unlock()

	// Copy keys if the value is map
	if src, ok := value.(mapContainer); ok {
		// Create the container
		if m.container == nil {
			m.container = make(mapContainer)
		}

		// Copy values
		for key := range src {
			m.container[key] = src[key]
		}
	}
}

func (m *BasicModel) PathModel(path []string) (value interface{}) {
	m.RLock()
	defer m.RUnlock()

	// Return nil if the container is not created
	if m.container == nil {
		return nil
	}

	// Find and return the value
	value = m.container
	for _, key := range path {
		if child, ok := value.(mapContainer); ok {
			value = child[key]
		} else {
			return nil
		}
	}

	return value
}

func (m *BasicModel) SetPathModel(path []string, value interface{}) {
	m.Lock()
	defer m.Unlock()

	// Create the container
	if m.container == nil {
		m.container = make(mapContainer)
	}

	// Navigate to the map described in the path
	parent := &m.container
	for i, key := range path {
		// Do not navigate further if we reached the latest item in the path
		if i == len(path)-1 {
			break
		}

		// Test the parent map has the item and the item is of map type
		if item, ok := (*parent)[key]; ok {
			if child, ok := item.(mapContainer); ok {
				parent = &child
				continue
			}
		}

		// Add new child if we reach here
		newChild := make(mapContainer)
		(*parent)[key] = newChild
		parent = &newChild
	}

	// Set the new value
	(*parent)[path[len(path)-1]] = value
}
