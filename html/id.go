package html

type Id string

func (id Id) Get() string {
	return string(id)
}

func (id *Id) Set(other string) {
	*id = Id(other)
}

func (id Id) Markup() string {
	if len(id) > 0 {
		return " id=\"" + string(id) + "\""
	}

	return ""
}
