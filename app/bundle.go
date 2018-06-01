package app

import "gopkg.in/workanator/vuego.v1/session"

type Bundle struct {
	SessionIdentifier session.Identifier
	StartScreen       Screener
}
