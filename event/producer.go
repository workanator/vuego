package event

import "context"

type Producer interface {
	Produce(buf []Event, ctx context.Context) (n int, err error)
}
