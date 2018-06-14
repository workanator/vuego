package view

// Flow allows to operate application views.
type Flow interface {
	Push(v *View) (top *View)
	Pop() (top *View)
	Replace(v *View) (top *View)
}
