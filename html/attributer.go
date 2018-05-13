package html

type Attributer interface {
	Markuper

	Has(name string) bool
	Get(name string) interface{}
	Set(name string, value interface{})
}
