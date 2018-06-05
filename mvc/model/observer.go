package model

// Observer allows to subscribe on model changes. Model changes are reported observer implementing Syncer.
type Observer interface {
	Observe(observer Syncer)
}
