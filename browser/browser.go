package browser

// Browser describes host application it capable operate with.
type Browser interface {
	Launcher

	Name() string
}
