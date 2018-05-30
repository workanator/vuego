package model

type GetSeter interface {
	Get(name string) (value interface{})
	Set(name string, value interface{})
}
