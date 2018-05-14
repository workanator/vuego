package html

import "fmt"

type Ider interface {
	fmt.Stringer
	Markuper

	Set(id string)
}
