package html

import "fmt"

const (
	Pixel Unit = iota
	Percentage
	Factor
)

// The unit of position or dimension.
type Unit uint8

// Format number with unit sign in CSS terms.
func (u Unit) Format(n int) string {
	switch u {
	case Factor:
		return fmt.Sprintf("%dpx", n*10)
	case Percentage:
		return fmt.Sprintf("%d%%", n)
	default:
		return fmt.Sprintf("%dpx", n)
	}
}
