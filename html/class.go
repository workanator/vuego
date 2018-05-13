package html

import "strings"

type Class []string

func (c *Class) Has(cls string) bool {
	if *c != nil {
		return c.find(cls) != -1
	}

	return false
}

func (c *Class) Add(cls string) {
	if *c == nil {
		*c = []string{cls}
	} else {
		if c.find(cls) == -1 {
			*c = append(*c, cls)
		}
	}
}

func (c *Class) Remove(cls string) {
	if *c != nil {
		if at := c.find(cls); at >= 0 {
			*c = append((*c)[:at], (*c)[at+1:]...)
		}
	}
}

func (c *Class) Clear() {
	if *c != nil {
		*c = (*c)[:0]
	}
}

func (c *Class) find(cls string) int {
	for index, item := range *c {
		if item == cls {
			return index
		}
	}

	return -1
}

func (c *Class) Markup() string {
	if *c == nil || len(*c) == 0 {
		return ""
	}

	return " class=\"" + strings.Join([]string(*c), " ") + "\""
}
