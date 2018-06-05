package browser

// Cannot find browser executable
type ErrBrowserNotFound struct{}

func (ErrBrowserNotFound) Error() string {
	return "browser is not found"
}

// Failed to launch browser
type ErrLaunchError string

func (e ErrLaunchError) Error() string {
	return string(e)
}
