package facade

import (
	"sync"

	"gopkg.in/workanator/vuego.v1/app/session"
)

// SessionPool implements session.Store interface and provides in-memoty session store based on sync.Map.
type SessionPool struct {
	New  session.Starter
	pool sync.Map
}

// Implement session.Starter interface. Start a new session with starter New.
func (sp *SessionPool) Start(sessId string) (sess *session.Session, err error) {
	// I think the better approach is to not check New for nil and to crash if it's nil instead of returning error.
	// Error returned can be ignored and that may lead to unpredicted behavior.
	sess, err = sp.New.Start(sessId)
	if err != nil {
		return sess, err
	}

	sp.pool.Store(sess.Id, sess)
	return sess, nil
}

// Implement session.Persister interface. Store the session sess in map.
func (sp *SessionPool) Persist(sess *session.Session) error {
	sp.pool.Store(sess.Id, sess)
	return nil
}

// Implement session.Restorer interface. Restore the session from map.
func (sp *SessionPool) Restore(sessId string) (sess *session.Session, err error) {
	if v, ok := sp.pool.Load(sessId); ok {
		return v.(*session.Session), nil
	}

	return nil, nil
}

// Implement session.Store interface. Test the session with the sessId exists in map.
func (sp *SessionPool) Exists(sessId string) bool {
	_, ok := sp.pool.Load(sessId)
	return ok
}
