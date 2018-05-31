package model

// Syncer is uniform interface for model synchronisation. Empty name means the whole model is being updated.
type Syncer interface {
	Sync(name string, value interface{})
}
