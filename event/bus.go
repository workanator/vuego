package event

// Bus connects event Consumer and Producer.
type Bus interface {
	Consumer
	Producer
}
