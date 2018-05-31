package html

// The error ErrRenderFailed denotes the situation when Rendering failed.
// With the field Reason the error can form a cascade structure which contains the full path to
// the element which failed to produce markup.
type ErrRenderFailed struct {
	Class  string
	Id     string
	Reason error
}

func (err ErrRenderFailed) Why() error {
	return err.Reason
}

func (err ErrRenderFailed) Error() string {
	var msg string

	// Write class
	if len(err.Class) > 0 {
		msg += err.Class
	} else {
		msg += "?"
	}

	// Write id
	if len(err.Id) > 0 {
		msg += "#" + err.Id
	}

	switch err.Reason.(type) {
	case ErrMarkupFailed, *ErrMarkupFailed:
		msg += "." + err.Reason.Error()
	case ErrRenderFailed, *ErrRenderFailed:
		msg += "." + err.Reason.Error()
	default:
		msg += " " + err.Reason.Error()
	}

	return msg
}
