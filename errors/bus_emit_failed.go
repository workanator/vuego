package errors

type ErrBusEmitFailed struct {
	Reason error
}

func (e ErrBusEmitFailed) Error() string {
	if e.Reason != nil {
		return "bus emit failed: " + e.Reason.Error()
	}

	return "bus emit failed"
}
