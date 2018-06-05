package event

type Producer interface {
	Produce(buf []Event) (n int, err error)
}
