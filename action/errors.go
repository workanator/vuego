package action

// The common action parsing error.
type ErrParse struct {
	Reason error
}

func (e ErrParse) Error() string {
	if e.Reason != nil {
		return "parse failed"
	} else {
		return "parse failed: " + e.Reason.Error()
	}
}

// Invalid command in action string.
type ErrInvalidCommand struct{}

func (ErrInvalidCommand) Error() string {
	return "invalid command"
}

// Invalid parenthesis in action string.
type ErrInvalidParenthesis struct{}

func (ErrInvalidParenthesis) Error() string {
	return "invalid parenthesis"
}
