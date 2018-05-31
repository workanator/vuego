package event

type Consumer interface {
	Consume(event Event) error
}
