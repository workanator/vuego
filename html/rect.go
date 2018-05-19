package html

const (
	bitRectTop = 1 << iota
	bitRectTopPercentage
	bitRectBottom
	bitRectBottomPercentage
	bitRectHeight
	bitRectHeightPercentage
	bitRectLeft
	bitRectLeftPercentage
	bitRectRight
	bitRectRightPercentage
	bitRectWidth
	bitRectWidthPercentage

	bitRectAll = bitRectTop | bitRectBottom | bitRectHeight | bitRectLeft | bitRectRight | bitRectWidth
)

type Rect struct {
	Top    int
	Bottom int
	Height int
	Left   int
	Right  int
	Width  int
	flags  uint16
}

func (rect *Rect) SetTop(top int, unit Unit) *Rect {
	rect.Top = top
	rect.flags |= bitRectTop

	if unit == Percentage {
		rect.flags |= bitRectTopPercentage
	} else {
		rect.flags = rect.flags &^ bitRectTopPercentage
	}

	return rect
}

func (rect Rect) WithTop(top int, unit Unit) Rect {
	return *rect.SetTop(top, unit)
}

func (rect *Rect) SetBottom(bottom int, unit Unit) *Rect {
	rect.Bottom = bottom
	rect.flags |= bitRectBottom

	if unit == Percentage {
		rect.flags |= bitRectBottomPercentage
	} else {
		rect.flags = rect.flags &^ bitRectBottomPercentage
	}

	return rect
}

func (rect Rect) WithBottom(bottom int, unit Unit) Rect {
	return *rect.SetBottom(bottom, unit)
}

func (rect *Rect) SetHeight(height int, unit Unit) *Rect {
	rect.Height = height
	rect.flags |= bitRectHeight

	if unit == Percentage {
		rect.flags |= bitRectHeightPercentage
	} else {
		rect.flags = rect.flags &^ bitRectHeightPercentage
	}

	return rect
}

func (rect Rect) WithHeight(height int, unit Unit) Rect {
	return *rect.SetHeight(height, unit)
}

func (rect *Rect) SetLeft(left int, unit Unit) *Rect {
	rect.Left = left
	rect.flags |= bitRectLeft

	if unit == Percentage {
		rect.flags |= bitRectLeftPercentage
	} else {
		rect.flags = rect.flags &^ bitRectLeftPercentage
	}

	return rect
}

func (rect Rect) WithLeft(left int, unit Unit) Rect {
	return *rect.SetLeft(left, unit)
}

func (rect *Rect) SetRight(right int, unit Unit) *Rect {
	rect.Right = right
	rect.flags |= bitRectRight

	if unit == Percentage {
		rect.flags |= bitRectRightPercentage
	} else {
		rect.flags = rect.flags &^ bitRectRightPercentage
	}

	return rect
}

func (rect Rect) WithRight(right int, unit Unit) Rect {
	return *rect.SetRight(right, unit)
}

func (rect *Rect) SetWidth(width int, unit Unit) *Rect {
	rect.Width = width
	rect.flags |= bitRectWidth

	if unit == Percentage {
		rect.flags |= bitRectWidthPercentage
	} else {
		rect.flags = rect.flags &^ bitRectWidthPercentage
	}

	return rect
}

func (rect Rect) WithWidth(width int, unit Unit) Rect {
	return *rect.SetWidth(width, unit)
}

func (rect *Rect) HasAny() bool {
	return rect.flags&bitRectAll != 0
}

func (rect *Rect) HasTop() bool {
	return rect.flags&bitRectTop == bitRectTop
}

func (rect *Rect) TopUnit() Unit {
	if rect.flags&bitRectTopPercentage == bitRectTopPercentage {
		return Percentage
	}

	return Pixel
}

func (rect *Rect) HasBottom() bool {
	return rect.flags&bitRectBottom == bitRectBottom
}

func (rect *Rect) BottomUnit() Unit {
	if rect.flags&bitRectBottomPercentage == bitRectBottomPercentage {
		return Percentage
	}

	return Pixel
}

func (rect *Rect) HasHeight() bool {
	return rect.flags&bitRectHeight == bitRectHeight
}

func (rect *Rect) HeightUnit() Unit {
	if rect.flags&bitRectHeightPercentage == bitRectHeightPercentage {
		return Percentage
	}

	return Pixel
}

func (rect *Rect) HasLeft() bool {
	return rect.flags&bitRectLeft == bitRectLeft
}

func (rect *Rect) LeftUnit() Unit {
	if rect.flags&bitRectLeftPercentage == bitRectLeftPercentage {
		return Percentage
	}

	return Pixel
}

func (rect *Rect) HasRight() bool {
	return rect.flags&bitRectRight == bitRectRight
}

func (rect *Rect) RightUnit() Unit {
	if rect.flags&bitRectRightPercentage == bitRectRightPercentage {
		return Percentage
	}

	return Pixel
}

func (rect *Rect) HasWidth() bool {
	return rect.flags&bitRectWidth == bitRectWidth
}

func (rect *Rect) WidthUnit() Unit {
	if rect.flags&bitRectWidthPercentage == bitRectWidthPercentage {
		return Percentage
	}

	return Pixel
}

func (rect *Rect) Impose(el *Element) {
	if el != nil {
		if rect.HasTop() {
			el.Style.Set("top", rect.TopUnit().Format(rect.Top))
		}

		if rect.HasBottom() {
			el.Style.Set("bottom", rect.BottomUnit().Format(rect.Bottom))
		}

		if rect.HasHeight() {
			el.Style.Set("height", rect.HeightUnit().Format(rect.Height))
		}

		if rect.HasLeft() {
			el.Style.Set("left", rect.LeftUnit().Format(rect.Left))
		}

		if rect.HasRight() {
			el.Style.Set("right", rect.RightUnit().Format(rect.Right))
		}

		if rect.HasWidth() {
			el.Style.Set("width", rect.WidthUnit().Format(rect.Width))
		}
	}
}
