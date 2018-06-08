package html

type Id string

func (id Id) IsEmpty() bool {
	return len(id) == 0
}

func (id Id) Equal(s string) bool {
	return string(id) == s
}

func (id Id) Markup() (string, error) {
	if len(id) > 0 {
		return " id=\"" + string(id) + "\"", nil
	}

	return "", nil
}

func (id Id) String() string {
	return string(id)
}
