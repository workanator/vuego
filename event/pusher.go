package event

type Pusher interface {
	Push(event *Event) error
}
