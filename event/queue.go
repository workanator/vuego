package event

import "fmt"

// Queue implements Producer interface and allows multiple producers to generate events for single bus.
// The implementation is thread-safe.
type Queue struct {
	queue chan Event
}

func NewQueue(size int) *Queue {
	if size > 1 {
		return &Queue{
			queue: make(chan Event, size),
		}
	}

	return &Queue{
		queue: make(chan Event),
	}
}

func (q *Queue) Push(e Event) {
	if q.queue != nil {
		q.queue <- e
	}
}

func (q *Queue) Produce(buf *[]Event) (n int, err error) {
	// Validate the buffer is not nil and have non-zero length
	if buf == nil {
		return 0, fmt.Errorf("buf is nil")
	} else if len(*buf) == 0 {
		return 0, fmt.Errorf("buf is zero size")
	}

	// Read the first event in blocking mode
	n = 1
	(*buf)[0] = <-q.queue

	// Read from the queue as much events as possible in non-blocking mode
	for n < len(*buf) {
		select {
		case e := <-q.queue:
			(*buf)[n] = e
			n++

		default:
			break
		}
	}

	return n, nil
}
