package event

type ErrEmitFailed struct {
	Reason error
}

func (e ErrEmitFailed) Error() string {
	if e.Reason != nil {
		return "bus emit failed: " + e.Reason.Error()
	}

	return "bus emit failed"
}
