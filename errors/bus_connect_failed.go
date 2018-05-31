package errors

type ErrBusConnectFailed struct {
	Reason error
}

func (e ErrBusConnectFailed) Error() string {
	if e.Reason != nil {
		return "bus connect failed: " + e.Reason.Error()
	}

	return "bus connect failed"
}
