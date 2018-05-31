package event

type ErrConsumeFailed struct {
	Reason error
}

func (e ErrConsumeFailed) Error() string {
	if e.Reason != nil {
		return "bus consume failed: " + e.Reason.Error()
	}

	return "bus consume failed"
}
