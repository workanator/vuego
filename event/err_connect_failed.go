package event

type ErrConnectFailed struct {
	Reason error
}

func (e ErrConnectFailed) Error() string {
	if e.Reason != nil {
		return "bus connect failed: " + e.Reason.Error()
	}

	return "bus connect failed"
}
