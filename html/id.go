package html

type Id string

func (id Id) IsEmpty() bool {
	return len(id) == 0
}

func (id Id) Markup() string {
	if len(id) > 0 {
		return " id=\"" + string(id) + "\""
	}

	return ""
}

func (id Id) String() string {
	return string(id)
}
