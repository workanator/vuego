package html

type Styler interface {
	Markuper

	Set(attr, style string)
	Setf(attr, styleFormat string, args ...interface{})
}
