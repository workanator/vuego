package errors

// The error ErrMarkupFailed denotes the situation when Markup rendering failed.
// With the field Reason the error can form a cascade structure which contains the full path to
// the element which failed to produce markup.
type ErrMarkupFailed struct {
	Tag    string
	Id     string
	Reason error
}

func (err ErrMarkupFailed) Error() string {
	var msg string

	// Write tag
	if len(err.Tag) > 0 {
		msg += err.Tag
	} else {
		msg += "?"
	}

	// Write id
	if len(err.Id) > 0 {
		msg += "#" + err.Id
	}

	switch v := err.Reason.(type) {
	case ErrMarkupFailed:
		msg += "." + v.Error()
	case *ErrMarkupFailed:
		msg += "." + v.Error()
	default:
		msg += " " + v.Error()
	}

	return msg
}
