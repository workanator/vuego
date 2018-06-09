package event

type Listener interface {
	Listen(buf []Event)
}
