package html

type Void struct{}

func (Void) Markup() (string, error) {
	return "", nil
}
