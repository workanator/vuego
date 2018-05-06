package browser

type Launcher interface {
	Launch(url string, options *LaunchOptions) error
}

type ErrBrowserNotFound struct{}

func (ErrBrowserNotFound) Error() string {
	return "browser is not found"
}
