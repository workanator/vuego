package event

type Puller interface {
	Pull() (event *Event, err error)
}
