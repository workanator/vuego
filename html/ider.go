package html

import "fmt"

type Ider interface {
	fmt.Stringer
	Markuper

	IsEmpty() bool
}
