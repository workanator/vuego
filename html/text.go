package html

import "html"

type Text string

func (t Text) Markup() (string, error) {
	return html.EscapeString(string(t)), nil
}
