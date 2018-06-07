package app

type Shutdowner interface {
	Shutdown(bundle *Bundle) error
}

// Wrapper to Shutdowner to make it posisble to invoke standalone function.
type ShutdownFunc func(bundle *Bundle) error

func (sf ShutdownFunc) Shutdown(bundle *Bundle) error {
	return sf(bundle)
}
