package server

type Refresher interface {
	Refresh() error
}
