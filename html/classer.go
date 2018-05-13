package html

type Classer interface {
	Markuper

	Has(cls string) bool
	Add(cls string)
	Remove(cls string)
	Clear()
}
