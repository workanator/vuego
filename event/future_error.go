package event

import (
	"sync"
	"sync/atomic"
)

type FutureError struct {
	lock     sync.Mutex
	awaiting uint32
	err      error
}

func NewFutureError() *FutureError {
	fe := &FutureError{
		awaiting: 1,
	}
	fe.lock.Lock()

	return fe
}

func (fe *FutureError) Complete(err error) {
	if atomic.CompareAndSwapUint32(&fe.awaiting, 1, 0) {
		fe.err = err
		fe.lock.Unlock()
	}
}

func (fe *FutureError) Wait() error {
	if atomic.LoadUint32(&fe.awaiting) == 1 {
		fe.lock.Lock()
		fe.lock.Unlock()
	}

	return fe.err
}

func (fe *FutureError) IsFinished() bool {
	return atomic.LoadUint32(&fe.awaiting) == 0
}
