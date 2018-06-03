package test_todo

import (
	"context"
	"net/http"
	"time"

	"gopkg.in/workanator/vuego.v1/event"
	"gopkg.in/workanator/vuego.v1/session"
)

type sessManager struct {
	sess *session.Session
}

func (sm *sessManager) Resolve(r *http.Request) (sess *session.Session, err error) {
	if sm.sess == nil {
		pushQue := event.NewPushQueue(64)
		pullQue := event.NewPullQueue(64)

		sm.sess = &session.Session{
			Context:        context.Background(),
			Id:             "1",
			User:           nil,
			State:          nil,
			InboundEvents:  &pushQue,
			OutboundEvents: &pullQue,
		}

		go func() {
			for {
				chan *event.Event(pullQue) <- &event.Event{
					Category: "data",
					Name:     "message",
					Time:     time.Now(),
					Data:     time.Now().String(),
				}

				time.Sleep(250 * time.Millisecond)
			}
		}()
	}

	return sm.sess, nil
}
