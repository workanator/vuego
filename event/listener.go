package event

type Listener interface {
	Listen(event Event)
}
