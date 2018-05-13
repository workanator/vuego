package html

import "strings"

type Multiple []Markuper

func (m Multiple) Markup() string {
	if len(m) == 0 {
		return ""
	}

	markup := strings.Builder{}
	for _, el := range m {
		markup.WriteString(el.Markup())
	}

	return markup.String()
}
