package html

const (
	OverflowInheritX Overflow = 0
	OverflowVisibleX Overflow = 1
	OverflowHiddenX  Overflow = 2
	OverflowScrollX  Overflow = 3
	OverflowAutoX    Overflow = 4
	overflowMaskX    Overflow = 7
)

const (
	OverflowInheritY Overflow = 0 << 3
	OverflowVisibleY Overflow = 1 << 3
	OverflowHiddenY  Overflow = 2 << 3
	OverflowScrollY  Overflow = 3 << 3
	OverflowAutoY    Overflow = 4 << 3
	overflowMaskY    Overflow = 7 << 3
)

const (
	OverflowInheritXY Overflow = OverflowInheritX | OverflowInheritY
	OverflowVisibleXY Overflow = OverflowVisibleX | OverflowVisibleY
	OverflowHiddenXY  Overflow = OverflowHiddenX | OverflowHiddenY
	OverflowScrollXY  Overflow = OverflowScrollX | OverflowScrollY
	OverflowAutoXY    Overflow = OverflowAutoX | OverflowAutoY
)

type Overflow uint8

func (o Overflow) X() Overflow {
	return o & overflowMaskX
}

func (o Overflow) Y() Overflow {
	return o & overflowMaskY
}

func (o Overflow) Impose(el *Element) {
	if el != nil {
		// Impose overflow X style
		switch o.X() {
		case OverflowInheritX:
			if el.Style.Has("overflow-x") {
				el.Style.Set("overflow-x", "inherit")
			}
		case OverflowVisibleX:
			el.Style.Set("overflow-x", "visible")
		case OverflowHiddenX:
			el.Style.Set("overflow-x", "hidden")
		case OverflowScrollX:
			el.Style.Set("overflow-x", "scroll")
		case OverflowAutoX:
			el.Style.Set("overflow-x", "auto")
		}

		// Impose overflow Y style
		switch o.Y() {
		case OverflowInheritY:
			if el.Style.Has("overflow-y") {
				el.Style.Set("overflow-y", "inherit")
			}
		case OverflowVisibleY:
			el.Style.Set("overflow-y", "visible")
		case OverflowHiddenY:
			el.Style.Set("overflow-y", "hidden")
		case OverflowScrollY:
			el.Style.Set("overflow-y", "scroll")
		case OverflowAutoY:
			el.Style.Set("overflow-y", "auto")
		}
	}
}
