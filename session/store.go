package session

type Store interface {
	Starter
	Persister
	Restorer

	Exists(sessId string) bool
}
