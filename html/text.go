package html

import "html"

type Text string

func (t Text) Markup() string {
	return html.EscapeString(string(t))
}
