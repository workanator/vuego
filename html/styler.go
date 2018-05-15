package html

type Styler interface {
	Markuper

	IsEmpty() bool
	Set(attr, style string)
	Setf(attr, styleFormat string, args ...interface{})
}
