package ui

const (
	OnAbort             Event = "onabort"
	OnBeforeInput             = "onbeforeinput"
	OnBlur                    = "onblur"
	OnClick                   = "onclick"
	OnCompositionEnd          = "oncompositionend"
	OnCompositionStart        = "oncompositionstart"
	OnCompositionUpdate       = "oncompositionupdate"
	OnDblClick                = "ondblclick"
	OnError                   = "onerror"
	OnFocus                   = "onfocus"
	OnFocusIn                 = "onfocusin"
	OnFocusOut                = "onfocusout"
	OnInput                   = "oninput"
	OnKeyDown                 = "onkeydown"
	OnKeyPress                = "onkeypress"
	OnKeyUp                   = "onkeyup"
	OnLoad                    = "onload"
	OnMouseDown               = "onmousedown"
	OnMouseEnter              = "onmouseenter"
	OnMouseLeave              = "onmouseleave"
	OnMouseMove               = "onmousemove"
	OnMouseOut                = "onmouseout"
	OnMouseOver               = "onmouseover"
	OnMouseUp                 = "onmouseup"
	OnResize                  = "onresize"
	OnScroll                  = "onscroll"
	OnSelect                  = "onselect"
	OnUnload                  = "onunload"
	OnWheel                   = "onwheel"
)

// The predefined set of possible events running in UI.
type Event string

func (e Event) String() string {
	return string(e)
}
