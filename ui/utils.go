package ui

import "gopkg.in/workanator/vuego.v1/html"

func ApplyToEl(el *html.Element, cmp Component, something interface{}) error {
	if el == nil || something == nil {
		return nil
	}

	// Impose
	switch v := something.(type) {
	case html.Imposer:
		v.Impose(el)

	case EventImposer:
		if err := v.ImposeEvent(el, cmp); err != nil {
			return err
		}
	}

	return nil
}
