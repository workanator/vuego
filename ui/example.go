package ui

/*

form := &Form{
	Id: "login_form",
	Title: "Sign In",
	Layout: layout.Vertical{},
	Items: []ui.Component{
		&ui.TextField{
			Id: "username",
			Label: "Login",
		},
		&ui.PasswordField{
			Id: "password",
			Label: "Password",
		},
		&ui.Button{
			Label: "Submit",
			OnClick: func(session *server.Session, cmp *ui.Component, event interface{}) {
				session.ReplaceScreen(myMainScreen())
			},
		},
	},
}

*/

/*
HTML

	e := &html.Element{
		Tag:       "div",
		Id:        "test",
		Class:     html.Class{"x-drop-shadow"},
		Style:     html.Style{"order": "hopper"},
		Attribute: html.Attribute{"ref": "_blank"},
		Inner: html.Multiple{
			&html.Element{
				Tag:       "label",
				Attribute: html.Attribute{"for": "login"},
				Inner:     html.Text("Login"),
			},
			&html.Element{
				Tag: "input",
				Attribute: html.Attribute{
					"name":  "login",
					"type":  "text",
					"value": "User",
				},
			},
		},
	}

	fmt.Println(e.Markup())
	return

*/

/*
COMPONENT

	e := &ui.Text{
		Tag: ui.Tag{
			Id: "plain_text",
		},
		Bounds: ui.Bounds{
			Rect:     ui.Rect{}.WithLeft(100).WithRight(200),
			Position: ui.PositionAbsolute,
			Overflow: ui.OverflowHiddenXY,
		},
		Text: "Hello World!",
		Type: ui.TextBlockquote,
	}

	fmt.Println(e.Render(nil, ui.Rect{}).Markup())
	return
*/
