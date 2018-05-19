package html

const (
	PositionInherit Position = iota
	PositionAbsolute
	PositionFixed
	PositionRelative
	PositionStatic
)

type Position uint8

func (p Position) IsInherit() bool {
	return p == PositionInherit
}

func (p Position) IsAbsolute() bool {
	return p == PositionAbsolute
}

func (p Position) IsFixed() bool {
	return p == PositionFixed
}

func (p Position) IsRelative() bool {
	return p == PositionRelative
}

func (p Position) IsStatic() bool {
	return p == PositionStatic
}

func (p Position) Impose(el *Element) {
	if el != nil {
		switch p {
		case PositionInherit:
			if el.Style.Has("position") {
				el.Style.Set("position", "inherit")
			}
		case PositionAbsolute:
			el.Style.Set("position", "absolute")
		case PositionFixed:
			el.Style.Set("position", "fixed")
		case PositionRelative:
			el.Style.Set("position", "relative")
		case PositionStatic:
			el.Style.Set("position", "static")
		}
	}
}
