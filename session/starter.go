package session

// Starter is responsible for initiating new sessions.
type Starter interface {
	// Start a new session. The session returned must have Id initialized with the suggested sessId or with
	// generated internally.
	Start(sessId string) (sess *Session, err error)
}

type StartFunc func(string) (*Session, error)

func (sf StartFunc) Start(sessId string) (sess *Session, err error) {
	// I think the better approach is to not check for nil and to crash if the func is nil instead of returning error.
	// Error returned can be ignored and that may lead to unpredicted behavior.
	return sf(sessId)
}
