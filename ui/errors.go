package ui

type ErrNil struct{}

func (ErrNil) Error() string {
	return "nil"
}
