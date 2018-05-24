package model

type PathModeler interface {
	PathModel(path []string) (value interface{})
	SetPathModel(path []string, value interface{})
}
