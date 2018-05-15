package html

type Styler interface {
	Markuper

	IsEmpty() bool
	Has(attr string) bool
	Set(attr, style string)
	Setf(attr, styleFormat string, args ...interface{})
}
