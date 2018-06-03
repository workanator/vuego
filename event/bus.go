package event

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const (
	eventBufSize = 64
)

const (
	BusDisconnected uint32 = iota
	BusConnecting
	BusConnected
	BusDiconnecting
)

type Bus struct {
	sync.RWMutex
	listeners []Listener
	state     uint32
	done      chan error
}

// Construct Bus instance. To make it working the method Connect should be used.
func NewBus() *Bus {
	// Construct and return bus
	return &Bus{
		listeners: make([]Listener, 0),
		state:     BusDisconnected,
	}
}

// Done returns the channel which block the current goroutine while the bus is connected and delivering
// events. When the bus is disconnected the channel will be closed what will unblock reading goroutine.
// In case of internal error it will be pushed in the channel and bus will be disconnected.
func (b *Bus) Done() <-chan error {
	return b.done
}

// Test if the bus is disconnected.
func (b *Bus) IsDisconnected() bool {
	return atomic.LoadUint32(&b.state) == BusDisconnected
}

// Connect starts listening emitter for new events. When events are available they are passed through listeners.
// At the end all events are passed to the consumer. And then events are discarded.
func (b *Bus) Connect(producer Producer, consumer Consumer) error {
	// Require the bus to be disconnected
	if !atomic.CompareAndSwapUint32(&b.state, BusDisconnected, BusConnecting) {
		return ErrConnectFailed{
			Reason: fmt.Errorf("bus connected"),
		}
	}

	// Reset to disconnected state at end
	defer atomic.StoreUint32(&b.state, BusDisconnected)

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

	// Create Done channel
	b.done = make(chan error)

	// Make buffer big enough for reading emitted events
	buf := make([]Event, eventBufSize)

	// Bus is ready to deliver events
	atomic.StoreUint32(&b.state, BusConnected)

	// Start and infinite loop of event delivery
	go func() {
		// Close Done channel at the end
		defer close(b.done)

		for {
			// Exit the loop if the bus in disconnecting state, i.e. Disconnect was invoked
			if atomic.LoadUint32(&b.state) == BusDiconnecting {
				break
			}

			// Ask the producer for new events
			n, err := producer.Produce(&buf)
			if err != nil {
				b.done <- ErrEmitFailed{
					Reason: err,
				}
				break
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
						b.done <- ErrConsumeFailed{
							Reason: err,
						}
						break
					}
				}
			}
		}
	}()

	return nil
}

// Disconnect the bus gently.
func (b *Bus) Disconnect() {
	atomic.CompareAndSwapUint32(&b.state, BusConnected, BusDiconnecting)
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
