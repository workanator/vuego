package session

import "net/http"

type Identifier interface {
	Identify(request *http.Request) (user *User, err error)
}
