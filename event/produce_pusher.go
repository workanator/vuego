package event

type ProducePusher interface {
	Producer
	Pusher
}
