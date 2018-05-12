package browser

type ErrBrowserNotFound struct{}

func (ErrBrowserNotFound) Error() string {
	return "browser is not found"
}

type ErrLaunchError string

func (e ErrLaunchError) Error() string {
	return string(e)
}
