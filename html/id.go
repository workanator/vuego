package html

type Id string

func (id *Id) Set(other string) {
	*id = Id(other)
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
