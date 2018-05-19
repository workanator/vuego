package vuetify

import (
	"gopkg.in/workanator/vuego.v1/html"
	"gopkg.in/workanator/vuego.v1/ui"
)

// Class of the Text.
type AppClass struct{}

func (AppClass) Class() string             { return "App" }
func (AppClass) ExtendedClass() ui.Classer { return ui.ComponentClass{} }

// Vuetify Application.
type App struct {
	ui.Tag
	ui.Children
	Dark bool
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
func (app *App) Render(parent *html.Element, viewport ui.Rect) *html.Element {
	if len(app.Tag.Id) == 0 {
		app.Tag.Id = "app"
	}

	// Create application component
	appEl := &html.Element{
		Tag: "v-app",
		Id:  app.Tag.Id + "-cmp",
	}

	if app.Dark {
		appEl.Attribute.Set("dark", true)
	}

	// Add childrens
	if app.Children.Len() > 0 {
		app.Children.Impose(appEl)
		appEl.Inner = app.Children.Render()
	}

	// Create application container element
	el := app.Tag.Element()
	el.Tag = "div"
	el.Inner = appEl

	return el
}
