package facade

import (
	"sync"

	"gopkg.in/workanator/vuego.v1/view"
)

// ViewStack implements view.Flow interface on top of stack.
type ViewStack struct {
	mut   sync.Mutex
	stack []*view.View
}

// Implement Push method. Push view on top of the stack and return the view that was on the top before.
func (vs *ViewStack) Push(v *view.View) (top *view.View) {
	vs.mut.Lock()
	defer vs.mut.Unlock()

	// Create the stack if it's empty
	if vs.stack == nil {
		vs.stack = make([]*view.View, 0)
	}

	// Remember the view on the top and push the new
	if len(vs.stack) > 0 {
		top = vs.stack[len(vs.stack)-1]
	}

	vs.stack = append(vs.stack, v)

	return top
}

// Implement Pop method. Pop the view from the stack' top.
func (vs *ViewStack) Pop() (top *view.View) {
	vs.mut.Lock()
	defer vs.mut.Unlock()

	// Return nil and do nothing if the stack is empty or nil
	if vs.stack == nil || len(vs.stack) == 0 {
		return nil
	}

	// Remove the item on the top and return it
	top = vs.stack[len(vs.stack)-1]
	vs.stack = vs.stack[:len(vs.stack)-1]

	return top
}

// Implement Replace method. Replace the view on the top of the stack. If the stack is empty the view is added.
func (vs *ViewStack) Replace(v *view.View) (top *view.View) {
	vs.mut.Lock()
	defer vs.mut.Unlock()

	// Create the stack if it's empty or nil and put the new view
	if vs.stack == nil {
		vs.stack = make([]*view.View, 0)
	}

	if len(vs.stack) == 0 {
		vs.stack = append(vs.stack, v)
		return nil
	}

	// Remember the top view and replace it with the new
	top = vs.stack[len(vs.stack)-1]
	vs.stack[len(vs.stack)-1] = v

	return top
}
