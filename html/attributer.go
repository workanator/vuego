package html

type Attributer interface {
	Markuper

	IsEmpty() bool
	Has(name string) bool
	Get(name string) interface{}
	Set(name string, value interface{})
}
