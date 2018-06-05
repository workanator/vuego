package event

type Consumer interface {
	Consume(buf []Event) error
}
