package session

import "net/http"

type Identifier interface {
	Identify(r *http.Request) (sessId string, err error)
}

type IdentifyFunc func(*http.Request) (string, error)

func (idf IdentifyFunc) Identify(r *http.Request) (sessId string, err error) {
	// I think the better approach is to not check for nil and to crash if the func is nil instead of returning error.
	// Error returned can be ignored and that may lead to unpredicted behavior.
	return idf(r)
}
