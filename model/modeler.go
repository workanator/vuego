package model

type Modeler interface {
	Model() (value interface{})
	SetModel(value interface{})
}
