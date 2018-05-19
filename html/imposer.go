package html

// Imposer applies attributes to HTML element. Attributes applied should be class, style, and other custom
// attributes except id and inner attributes and except tag name.
type Imposer interface {
	// Impose attributes to the HTML element el.
	Impose(el *Element)
}
