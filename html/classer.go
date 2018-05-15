package html

type Classer interface {
	Markuper

	IsEmpty() bool
	Has(cls string) bool
	Add(cls string)
	Remove(cls string)
	Clear()
}
