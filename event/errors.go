package event

// Bus connected consumer failed to consume events
type ErrConsumeFailed struct {
	Reason error
}

func (e ErrConsumeFailed) Error() string {
	if e.Reason != nil {
		return "consume failed: " + e.Reason.Error()
	}

	return "consume failed"
}

// Bus connected provider failed to provide events
type ErrProduceFailed struct {
	Reason error
}

func (e ErrProduceFailed) Error() string {
	if e.Reason != nil {
		return "produce failed: " + e.Reason.Error()
	}

	return "produce failed"
}
