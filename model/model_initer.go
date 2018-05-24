package model

type ModelIniter interface {
	Modeler

	Init(value interface{}) Modeler
}
