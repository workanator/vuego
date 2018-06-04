package session

type Restorer interface {
	Restore(sessId string) (sess *Session, err error)
}
