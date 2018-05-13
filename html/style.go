package html

import (
	"fmt"
	"strings"
)

type Style map[string]string

func (s *Style) Set(attr, style string) {
	if *s == nil {
		*s = make(map[string]string)
	}

	(*s)[attr] = style
}

func (s *Style) Setf(attr, styleFormat string, args ...interface{}) {
	if *s == nil {
		*s = make(map[string]string)
	}

	(*s)[attr] = fmt.Sprintf(styleFormat, args...)
}

func (s *Style) Markup() string {
	if *s == nil || len(*s) == 0 {
		return ""
	}

	markup := strings.Builder{}
	for k := range *s {
		if markup.Len() > 0 {
			markup.WriteRune(';')
		}

		markup.WriteString(k)
		markup.WriteRune(':')
		markup.WriteString((*s)[k])
	}

	return " style=\"" + markup.String() + "\""
}
