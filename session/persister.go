package session

type Persister interface {
	Persist(sess *Session) error
}
