package errors

type ErrBusConsumeFailed struct {
	Reason error
}

func (e ErrBusConsumeFailed) Error() string {
	if e.Reason != nil {
		return "bus consume failed: " + e.Reason.Error()
	}

	return "bus consume failed"
}
