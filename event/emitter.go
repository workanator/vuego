package event

type Emitter interface {
	Emit(buf *[]Event) (n int, err error)
}
