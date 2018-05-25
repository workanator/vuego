package html

import (
	"strings"
)

type Multiple []Markuper

func (m Multiple) Markup() (string, error) {
	if len(m) == 0 {
		return "", nil
	}

	markup := strings.Builder{}
	for _, el := range m {
		if s, err := el.Markup(); err != nil {
			return "", err
		} else {
			markup.WriteString(s)
		}
	}

	return markup.String(), nil
}
