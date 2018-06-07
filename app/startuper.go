package app

type Startuper interface {
	Startup(bundle *Bundle) error
}

// Wrapper for Startuper to make it possible to invoke standalone function.
type StartupFunc func(bundle *Bundle) error

func (sf StartupFunc) Startup(bundle *Bundle) error {
	return sf(bundle)
}
