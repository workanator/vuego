package view

// Find is not found.
type ErrNotFound struct{}

func (ErrNotFound) Error() string {
	return "not found"
}
