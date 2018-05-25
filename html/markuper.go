package html

type Markuper interface {
	Markup() (markup string, err error)
}
