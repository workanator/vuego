package html

import (
	"fmt"
	"strings"
)

type Attribute map[string]interface{}

func (a *Attribute) IsEmpty() bool {
	return len(*a) == 0
}

func (a *Attribute) Has(name string) bool {
	if *a != nil {
		if _, ok := (*a)[name]; ok {
			return true
		}
	}

	return false
}

func (a *Attribute) Get(name string) interface{} {
	if *a != nil {
		return (*a)[name]
	}

	return nil
}

func (a *Attribute) Set(name string, value interface{}) {
	if *a == nil {
		*a = make(map[string]interface{})
	}

	(*a)[name] = value
}

func (a *Attribute) Markup() (string, error) {
	if *a == nil {
		return "", nil
	}

	markup := strings.Builder{}
	for k := range *a {
		v := (*a)[k]
		switch t := v.(type) {
		case bool:
			if t {
				markup.WriteRune(' ')
				markup.WriteString(k)
			}
		default:
			markup.WriteRune(' ')
			markup.WriteString(k)
			markup.WriteString("=\"")
			markup.WriteString(fmt.Sprint(v))
			markup.WriteRune('"')
		}
	}

	return markup.String(), nil
}
