package event

import "fmt"

// PushQueue implements Producer interface and allows multiple producers to generate events for single bus.
// The implementation is thread-safe.
type PushQueue chan *Event

func NewPushQueue(size int) PushQueue {
	if size > 1 {
		return make(chan *Event, size)
	}

	return make(chan *Event)
}

func (q *PushQueue) Close() error {
	if q != nil {
		close(*q)
	}

	return nil
}

func (q *PushQueue) Queue() chan<- *Event {
	return *q
}

func (q *PushQueue) Push(e *Event) error {
	if q != nil {
		*q <- e
	}

	return nil
}

func (q *PushQueue) Produce(buf *[]Event) (n int, err error) {
	// Validate the channel is created
	if q == nil {
		return 0, nil
	}

	// Validate the buffer is not nil and have non-zero length
	if buf == nil {
		return 0, fmt.Errorf("buf is nil") // TODO: better error
	} else if len(*buf) == 0 {
		return 0, fmt.Errorf("buf is zero size") // TODO: better error
	}

	// Read the first event in blocking mode
	n = 1
	(*buf)[0] = *<-*q

	// Read from the queue as much events as possible in non-blocking mode
	for n < len(*buf) {
		select {
		case e := <-*q:
			(*buf)[n] = *e
			n++

		default:
			break
		}
	}

	return n, nil
}
