package facade

import (
	"context"
	"crypto/md5"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"gopkg.in/workanator/vuego.v1/event"
	"gopkg.in/workanator/vuego.v1/session"
)

// SingleSessions supports only one session (single user approach). The session uses sync.RWMutex as data races
// protection and it is recommended to use the mutex embedded when accessing fields of the session.
type SingleSession struct {
	sync.RWMutex

	// The default context to supply the created session with. Will be initialized with context.Background()
	// if not set.
	Context context.Context

	// The default session identifier. The new random identifier will be generated when session is created
	// if the identifier is empty.
	Id string

	// The user instance the session contains.
	User *session.User

	// The state instance the session contains.
	State interface{}

	// Session event bus.
	EventBus event.Bus

	sess *session.Session
}

// Implement session.Resolver interface. The session is created the first time the resolution happens and
// then reused each time the resolution is issued.
func (ss *SingleSession) Resolve(r *http.Request) (*session.Session, error) {
	ss.Lock()
	defer ss.Unlock()

	// Create the new session when it issued for the first time.
	if ss.sess == nil {
		// Reused or create a new context
		if ss.Context == nil {
			ss.Context = context.Background()
		}

		// Generate the new session ID if it's empty
		if len(ss.Id) == 0 {
			v := fmt.Sprintf("%d.%d.%d.%d", os.Getgid(), os.Getpid(), time.Now().Unix(), rand.Int())
			ss.Id = fmt.Sprintf("%x", md5.Sum([]byte(v)))
		}

		// Create the session
		ss.sess = &session.Session{
			Context:  ss.Context,
			Id:       ss.Id,
			User:     ss.User,
			Data:     ss.State,
			EventBus: ss.EventBus,
		}
	}

	return ss.sess, nil
}
