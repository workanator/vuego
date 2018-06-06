package facade

import (
	"context"

	"gopkg.in/workanator/vuego.v1/app/event"
)

const (
	eventQueueSize = 64
)

// Reactor implements event.Bus interface and establishes event processing.
type Reactor struct {
	outEvents chan event.Event
}

func NewReactor() *Reactor {
	return &Reactor{
		outEvents: make(chan event.Event, eventQueueSize),
	}
}

// Implement event.Consumer interface. The enter point of inbound events.
func (r *Reactor) Consume(buf []event.Event, ctx context.Context) error {
	// Marshal events
	for _, ev := range buf {
		// Deliver system events directly to outbound queue
		if ev.Category.IsSystem() {
			r.outEvents <- ev
		}

		// Test if context is terminated
		if ctx != nil {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
		}
	}

	return nil
}

// Implement event.Producer interface. The leave point of outbound events.
func (r *Reactor) Produce(buf []event.Event, ctx context.Context) (n int, err error) {
	// Return immediately if buf is of zero length
	if len(buf) == 0 {
		return 0, nil
	}

	// Read the first event in blocking mode
	if ctx != nil {
		select {
		case buf[n] = <-r.outEvents:
			n++
		case <-ctx.Done():
			return 0, ctx.Err()
		}
	} else {
		buf[n] = <-r.outEvents
		n++
	}

	// Read as much events as possible from the queue
	hasMore := true
	for n < len(buf) && hasMore {
		if ctx != nil {
			select {
			case buf[n] = <-r.outEvents:
				n++
			case <-ctx.Done():
				return n, ctx.Err()
			default:
				hasMore = false
			}
		} else {
			select {
			case buf[n] = <-r.outEvents:
				n++
			default:
				hasMore = false
			}
		}
	}

	return n, nil
}
