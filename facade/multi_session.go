package facade

import (
	"net/http"

	"context"

	"gopkg.in/workanator/vuego.v1/event"
	"gopkg.in/workanator/vuego.v1/session"
)

// MultiSessions supports multiple sessions with persistence mechanism provided by session Store. By default
// newly created session' context is initialized with context.Background() unless session context fields is set.
type MultiSession struct {
	// Session identifier implementation.
	Who session.Identifier

	// Session store implementation.
	Store session.Store

	// Events producer. Inbound events go from client to server.
	InboundEvents event.Pusher

	// Events consumer. Outbound events go from server to client.
	OutboundEvents event.Puller
}

// Implement session.Resolver interface.
func (ms *MultiSession) Resolve(r *http.Request) (sess *session.Session, err error) {
	var sessId string

	// Identify the request first
	sessId, err = ms.Who.Identify(r)
	if err != nil {
		return nil, err
	}

	// Check if the session exists in the store
	if ms.Store.Exists(sessId) {
		sess, err = ms.Store.Restore(sessId)
	} else {
		// Assume the Store persist newly created sessions
		sess, err = ms.Store.Start(sessId)
	}

	// Setup uninitialized fields
	if sess != nil {
		if sess.Context == nil {
			sess.Context = context.Background()
		}

		if sess.InboundEvents == nil {
			sess.InboundEvents = ms.InboundEvents
		}

		if sess.OutboundEvents == nil {
			sess.OutboundEvents = ms.OutboundEvents
		}
	}

	return sess, err
}
