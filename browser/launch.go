package browser

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

func Lauch(url string, options *Options, browsers ...Browser) error {
	errs := make([]string, 0)
	for _, launcher := range browsers {
		logrus.WithField("browser", launcher.Name()).Info("Launch browser")

		if err := launcher.Launch(url, options); err != nil {
			if _, ok := err.(ErrBrowserNotFound); !ok {
				errs = append(errs, fmt.Sprintf("%s: %s", launcher.Name(), err.Error()))
			}
		} else {
			return nil
		}
	}

	if len(errs) > 0 {
		return ErrLaunchError(fmt.Sprintf(strings.Join(errs, "; ")))
	}

	return ErrLaunchError("cannot launch browser")
}
