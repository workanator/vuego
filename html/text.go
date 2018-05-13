package html

type Text string

func (t Text) Markup() string {
	return string(t)
}
