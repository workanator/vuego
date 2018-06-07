package app

type Lifecycle interface {
	Startuper
	Shutdowner
}

// Wrapper to Lifecycle to make it possible to invoke single function for any state change.
type LifecycleFunc func(starting bool, bundle *Bundle) error

func (lf LifecycleFunc) Startup(bundle *Bundle) error {
	return lf(true, bundle)
}

func (lf LifecycleFunc) Shutdown(bundle *Bundle) error {
	return lf(false, bundle)
}
