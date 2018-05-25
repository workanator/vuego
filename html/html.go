package html

type Html string

func (h Html) Markup() (string, error) {
	return string(h), nil
}
