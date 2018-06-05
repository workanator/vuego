package event

import "context"

type Consumer interface {
	Consume(buf []Event, ctx context.Context) error
}
