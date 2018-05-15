package ui

const (
	bitRectTop         = 1
	bitRectBottom      = 2
	bitRectLeft        = 4
	bitRectRight       = 8
	bitRectAll         = bitRectTop | bitRectBottom | bitRectLeft | bitRectRight
	bitRectFixedHeight = bitRectTop | bitRectBottom
	bitRectFixedWidth  = bitRectLeft | bitRectRight
)

type Rect struct {
	Top    int
	Bottom int
	Left   int
	Right  int
	flags  uint8
}

func (rect *Rect) SetTop(top int) *Rect {
	rect.Top = top
	rect.flags |= bitRectTop
	return rect
}

func (rect Rect) WithTop(top int) Rect {
	rect.Top = top
	rect.flags |= bitRectTop
	return rect
}

func (rect *Rect) SetBottom(bottom int) *Rect {
	rect.Bottom = bottom
	rect.flags |= bitRectBottom
	return rect
}

func (rect Rect) WithBottom(bottom int) Rect {
	rect.Bottom = bottom
	rect.flags |= bitRectBottom
	return rect
}

func (rect *Rect) SetLeft(left int) *Rect {
	rect.Left = left
	rect.flags |= bitRectLeft
	return rect
}

func (rect Rect) WithLeft(left int) Rect {
	rect.Left = left
	rect.flags |= bitRectLeft
	return rect
}

func (rect *Rect) SetRight(right int) *Rect {
	rect.Right = right
	rect.flags |= bitRectRight
	return rect
}

func (rect Rect) WithRight(right int) Rect {
	rect.Right = right
	rect.flags |= bitRectRight
	return rect
}

func (rect *Rect) HasAny() bool {
	return rect.flags&bitRectAll != 0
}

func (rect *Rect) HasTop() bool {
	return rect.flags&bitRectTop == bitRectTop
}

func (rect *Rect) HasBottom() bool {
	return rect.flags&bitRectBottom == bitRectBottom
}

func (rect *Rect) HasLeft() bool {
	return rect.flags&bitRectLeft == bitRectLeft
}

func (rect *Rect) HasRight() bool {
	return rect.flags&bitRectRight == bitRectRight
}

func (rect *Rect) IsFixedHeight() bool {
	return rect.flags&bitRectFixedHeight == bitRectFixedHeight
}

func (rect *Rect) IsFixedWidth() bool {
	return rect.flags&bitRectFixedWidth == bitRectFixedWidth
}

func (rect *Rect) Height() int {
	return rect.Bottom - rect.Top + 1
}

func (rect *Rect) Width() int {
	return rect.Right - rect.Left + 1
}
