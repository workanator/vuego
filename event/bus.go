package event

import (
	"fmt"
	"sync"
)

const (
	eventBufSize = 64
)

type Bus struct {
	sync.RWMutex
	listeners []Listener
}

// Construct Bus instance. To make it working the method Connect should be used.
func NewBus() *Bus {
	// Construct and return bus
	return &Bus{
		listeners: make([]Listener, 0),
	}
}

// Connect starts listening emitter for new events. When events are available they are passed through listeners.
// At the end all events are passed to the consumer. And then events are discarded.
// The method blocks the current goroutine.
func (b *Bus) Connect(producer Producer, consumer Consumer) error {
	// Validate producer and consumer
	if producer == nil {
		return ErrConnectFailed{
			Reason: fmt.Errorf("producer is nil"),
		}
	}

	if consumer == nil {
		return ErrConnectFailed{
			Reason: fmt.Errorf("consumer is nil"),
		}
	}

	// Make buffer big enough for reading emitted events
	buf := make([]Event, eventBufSize)

	// Start and infinite loop of event delivery
	for {
		// Ask the producer for new events
		n, err := producer.Produce(&buf)
		if err != nil {
			return ErrEmitFailed{
				Reason: err,
			}
		}

		// Deliver events if any
		if n > 0 {
			// Run emitted events through listeners
			func() {
				b.RLock()
				defer b.RUnlock()

				for i := 0; i < n; i++ {
					for _, listener := range b.listeners {
						listener.Listen(buf[i])
					}
				}
			}()

			// Push emitted events to consumer
			for i := 0; i < n; i++ {
				if err := consumer.Consume(buf[i]); err != nil {
					return ErrConsumeFailed{
						Reason: err,
					}
				}
			}
		}
	}

	return nil
}

// Attach event listener to the bus.
func (b *Bus) AttachListener(listener Listener) {
	if listener != nil {
		b.Lock()
		defer b.Unlock()

		if b.listeners == nil {
			b.listeners = make([]Listener, 0)
		}

		b.listeners = append(b.listeners, listener)
	}
}

// Detach event listener from the bus.
func (b *Bus) DetachListener(listener Listener) {
	if listener != nil {
		b.Lock()
		defer b.Unlock()

		if b.listeners != nil {
			// Find the listener
			var at int = -1
			for i := 0; i < len(b.listeners); i++ {
				if b.listeners[i] == listener {
					at = i
					break
				}
			}

			// Remove the listener if it's found
			if at >= 0 {
				b.listeners = append(b.listeners[:at], b.listeners[at+1:]...)
			}
		}
	}
}
