package browser

import (
	"fmt"
	"os/exec"
)

type firefox struct{}

func Firefox() Launcher {
	return &firefox{}
}

func (firefox) Launch(url string, options *LaunchOptions) error {
	// Prepare browser command line arguments
	args := make([]string, 0)

	// Create new instance of browser
	if options.NewInstance {
		args = append(args, "--new-instance")
	}

	// Configure browser window
	if options.Window != nil {
		// Window size
		if options.Window.Size != nil {
			args = append(args,
				"--window-size",
				fmt.Sprintf("%d,%d", options.Window.Size.Width, options.Window.Size.Height),
			)
		}
	}

	// Open URL in new window
	args = append(args, url)

	return exec.Command("firefox", args...).Run()
}
