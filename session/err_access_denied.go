package session

type ErrAccessDenied struct{}

func (ErrAccessDenied) Error() string {
	return "access denied"
}

// Test if the error is Access Denied
func IsAccessDenied(e error) bool {
	_, ok := e.(ErrAccessDenied)
	return ok
}
