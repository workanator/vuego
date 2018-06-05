package model

// ModelInitialer wraps the other interface implementing Modeler and adds a new method
// for setting the initial value.
type ModelInitialer interface {
	Modeler

	// Set the initial value in the owner model.
	Initial(value interface{}) Modeler
}

// The simple implementation of ModelInitialer which simply invokes SetModel method
// of the wrapped model.
type ModelInitial struct {
	Modeler
}

func (mi ModelInitial) Initial(value interface{}) Modeler {
	if mi.Modeler != nil {
		mi.Modeler.SetModel(value)
	}

	return mi.Modeler
}
