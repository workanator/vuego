package session

import "net/http"

type Resolver interface {
	Resolve(r *http.Request) (sess *Session, err error)
}

type ResolveFunc func(*http.Request) (*Session, error)

func (rf ResolveFunc) Resolve(r *http.Request) (sess *Session, err error) {
	// I think the better approach is to not check for nil and to crash if the func is nil instead of returning error.
	// Error returned can be ignored and that may lead to unpredicted behavior.
	return rf(r)
}
