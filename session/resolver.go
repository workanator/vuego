package session

import "net/http"

type Resolver interface {
	Resolve(r *http.Request) (sess *Session, err error)
}
