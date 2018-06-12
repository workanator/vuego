package server

type WriteFunc func(b []byte) (int, error)

func (wf WriteFunc) Write(b []byte) (int, error) {
	return wf(b)
}
