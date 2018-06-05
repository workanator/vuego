package model

type Propertier interface {
	Property(name string) (value interface{})
	SetProperty(name string, value interface{})
}
