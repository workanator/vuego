package browser

// Launcher is responsible for starting application.
type Launcher interface {
	Launch(url string, options *Options) error
}
