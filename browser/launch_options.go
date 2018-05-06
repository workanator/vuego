package browser

import "gopkg.in/workanator/vuego.v1/ui"

type LaunchOptions struct {
	Priority    []Launcher
	NewInstance bool
	Window      *ui.WindowOptions
}
