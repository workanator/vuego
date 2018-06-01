package app

type ScreenManager interface {
	PushScreen(screen Screener) (previos Screener)
	PopScreen(screen Screener) (previos Screener)
	ReplaceScreen(screen Screener) (previos Screener)
}
