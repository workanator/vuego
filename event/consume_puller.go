package event

type ConsumePuller interface {
	Consumer
	Puller
}
