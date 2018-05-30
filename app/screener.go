package app

import "gopkg.in/workanator/vuego.v1/html"

type Screener interface {
	// Get internal screen name.
	Name() string

	// Get user-friendly title.
	Title() string

	// Get head tags markuper.
	Head() html.Markuper

	// Get body renderer.
	Body() html.Renderer

	// Get models used on the screen.
	Models() []html.Markuper
}
