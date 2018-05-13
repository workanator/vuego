package html

type Ider interface {
	Markuper

	Get() string
	Set(id string)
}
