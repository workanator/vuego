package browser

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
)

type firefox struct{}

func Firefox() Browser {
	return &firefox{}
}

func (firefox) Name() string {
	return "Firefox"
}

func (firefox) Launch(url string, options *Options) error {
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
	args = append(args, "--new-tab", url)

	// Execute the command
	logrus.WithFields(logrus.Fields{
		"bin":  "firefox",
		"args": args,
	}).Info("Exec")

	if err := exec.Command("firefox", args...).Start(); err != nil {
		if os.IsNotExist(err) {
			return ErrBrowserNotFound{}
		}

		return err
	}

	return nil
}
