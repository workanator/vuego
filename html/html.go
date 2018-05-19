package html

type Html string

func (h Html) Markup() string {
	return string(h)
}
