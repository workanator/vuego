package vuetify

import (
	"gopkg.in/workanator/vuego.v1/html"
	"gopkg.in/workanator/vuego.v1/ui"
)

// Class of the App.
type AppClass struct{}

func (AppClass) Class() string             { return "App" }
func (AppClass) ExtendedClass() ui.Classer { return ui.ComponentClass{} }

// Vuetify Application.
type App struct {
	ui.Tag
	ui.Children
	Appearance Appearance
}

// Get class name.
func (App) Class() string {
	return AppClass{}.Class()
}

// Get the class name that class extends.
func (App) ExtendedClass() ui.Classer {
	return AppClass{}
}

// Render content into HTML Element.
func (app *App) Render(parent *html.Element, viewport html.Rect) *html.Element {
	if len(app.Tag.Id) == 0 {
		app.Tag.Id = "app"
	}

	// Create application component
	appEl := &html.Element{
		Tag: "v-app",
		Id:  app.Tag.Id + "-cmp",
	}

	if app.Appearance != Default {
		appEl.Attribute.Set(app.Appearance.String(), true)
	}

	// Add children items
	if app.Children.Len() > 0 {
		app.Children.Impose(appEl)
		appEl.Inner = app.Children.Render(appEl, viewport)
	}

	// Create application container element
	el := app.Tag.Element()
	el.Tag = "div"
	el.Inner = appEl

	return el
}
