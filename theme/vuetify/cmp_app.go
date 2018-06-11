package vuetify

import (
	"gopkg.in/workanator/vuego.v1/html"
	"gopkg.in/workanator/vuego.v1/model"
	"gopkg.in/workanator/vuego.v1/session"
	"gopkg.in/workanator/vuego.v1/ui"
	"gopkg.in/workanator/vuego.v1/vue"
)

// Class of the App.
type AppClass struct{}

func (AppClass) Class() string             { return "App" }
func (AppClass) ExtendedClass() ui.Classer { return ui.ElClass{} }

// Vuetify Application.
type App struct {
	ui.Tag
	Children   ui.ItemRenderer
	Appearance Appearance
	Toolbar    struct {
		Top    *Toolbar
		Bottom *Toolbar
	}
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
func (app *App) Render(parent *html.Element, viewport html.Rect) (*html.Element, error) {
	if len(app.Tag.Id) == 0 {
		app.Tag.Id = "app"
	}

	// Create content container
	contentEl := &html.Element{
		Tag: "v-content",
	}

	// Add children items
	if app.Children != nil {
		if el, err := app.Children.Render(contentEl, viewport); err != nil {
			return nil, html.ErrRenderFailed{
				Class:  app.Class(),
				Id:     app.Tag.Id.String(),
				Reason: err,
			}
		} else {
			contentEl.Inner = el
		}
	}

	// Create application component
	appEl := &html.Element{
		Tag: "v-app",
		Id:  app.Tag.Id + "-cmp",
	}

	app.Appearance.Impose(appEl)

	// Make application layout
	if app.Toolbar.Top == nil && app.Toolbar.Bottom == nil {
		appEl.Inner = contentEl
	} else {
		items := make([]html.Markuper, 0, 3)

		if app.Toolbar.Top != nil {
			if tbEl, err := app.Toolbar.Top.Render(appEl, viewport); err != nil {
				return nil, html.ErrRenderFailed{
					Class:  app.Class(),
					Id:     app.Tag.Id.String(),
					Reason: err,
				}
			} else {
				tbEl.Attribute.Set("app", true)
				items = append(items, tbEl)
			}
		}

		items = append(items, contentEl)

		if app.Toolbar.Bottom != nil {
			if tbEl, err := app.Toolbar.Bottom.Render(appEl, viewport); err != nil {
				return nil, html.ErrRenderFailed{
					Class:  app.Class(),
					Id:     app.Tag.Id.String(),
					Reason: err,
				}
			} else {
				tbEl.Attribute.Set("app", true)
				items = append(items, tbEl)
			}
		}

		appEl.Inner = html.Multiple(items)
	}

	// Follow with script tag with Vue instance
	// TODO: Should that be automated?
	scriptEl := &html.Element{
		Tag: "script",
		Inner: html.Multiple{
			html.Text("Vuego.App = "),
			&vue.Vue{
				Id:   app.Tag.Id,
				Data: &model.Container{},
			},
		},
	}

	// Create application container element
	el := app.Tag.Element()
	el.Tag = "div"
	el.Inner = appEl
	el.Chain = scriptEl

	return el, nil
}

// Implement Component interface.
func (app *App) MarshalEvent(sess *session.Session, target, event string, data interface{}) (processed bool, err error) {
	return processed, err
}
