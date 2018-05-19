package html

type Renderer interface {
	// Render content into HTML Element.
	Render(parent *Element, viewport Rect) *Element
}
