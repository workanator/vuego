package event

// PullQueue implements Consumer interface and allows multiple consumers to read events from single bus.
// The implementation is thread-safe.
type PullQueue chan *Event

func NewPullQueue(size int) PullQueue {
	if size > 1 {
		return make(chan *Event, size)
	}

	return make(chan *Event)
}

func (q *PullQueue) Close() error {
	if q != nil {
		close(*q)
	}

	return nil
}

func (q *PullQueue) Queue() <-chan *Event {
	return *q
}

func (q *PullQueue) Pull() (*Event, error) {
	if q != nil {
		return <-*q, nil
	}

	return nil, nil
}

func (q *PullQueue) Consume(e Event) (err error) {
	// Validate the channel is created
	if q == nil {
		return nil
	}

	*q <- &e
	return nil
}
