package browser

func Lauch(url string, options *LaunchOptions) error {
	for _, launcher := range options.Priority {
		if err := launcher.Launch(url, options); err != nil {
			if _, ok := err.(ErrBrowserNotFound); !ok {
				// TODO: bad
			}
		} else {
			break
		}
	}

	return nil
}
