package action

import (
	"strings"
)

// Action is a command w/ or w/o arguments to perform. Actions are encoded in function call syntax with
// function name, parenthesis, and argument list which can be empty.
// E.g.
// - Open(this/page)
// - Logout()
type Action struct {
	source  string
	command bounds
	args    []bounds
}

// Parse string into Action struct and return parsing error if any.
func Parse(s string) (act *Action, err error) {
	if len(s) == 0 {
		return nil, nil
	}

	act = &Action{
		source: s,
	}

	// First token should be command which ends at the end of the string
	// or when parenthesis starts
	pos := 1
	act.command.l = 0
	for pos < len(s) {
		if s[pos] == '(' {
			act.command.r = pos
			break
		} else {
			pos++
		}
	}

	if pos >= len(s) {
		return act, nil
	}

	// Test for empty parenthesis
	pos++
	if pos >= len(s) {
		return act, ErrParse{
			Reason: ErrInvalidParenthesis{},
		}
	} else if s[pos] == ')' {
		// Stop parsing after closing bracket
		return act, nil
	}

	// Parse argument list
	act.args = make([]bounds, 0)

	var arg bounds
	arg.l = pos
	pos++
	for pos < len(s) {
		if s[pos] == ',' {
			arg.r = pos
			act.args = append(act.args, arg)
			arg = emptyBounds()
		} else if s[pos] == ')' {
			// Stop parsing after closing bracket
			arg.r = pos
			act.args = append(act.args, arg)
			break
		} else if pos == len(s)-1 {
			return act, ErrParse{
				Reason: ErrInvalidParenthesis{},
			}
		}

		pos++
	}

	return act, nil
}

// Construct action instance.
func New(command string, args ...string) *Action {
	if len(args) > 0 {
		act := command + "(" + strings.Join(args, ",") + ")"
		action, _ := Parse(act)
		return action
	}

	return &Action{
		source: command + "()",
		command: bounds{
			l: 0,
			r: len(command) - 1,
		},
	}
}

// Get source string of the action from which it was parsed.
func (a *Action) Source() string {
	return a.source
}

// Get command if the action is valid.
func (a *Action) Command() string {
	if a.command.IsValid() {
		return a.source[a.command.l:a.command.r]
	}

	return ""
}

// Get argument count.
func (a *Action) ArgCount() int {
	return len(a.args)
}

// Get n-th argument.
func (a *Action) Arg(n int) string {
	if n >= 0 && n < len(a.args) {
		if a.args[n].IsValid() {
			return a.source[a.args[n].l:a.args[n].r]
		}
	}

	return ""
}

// Get JS representation of action.
func (a *Action) JS() string {
	s := strings.Builder{}
	s.WriteString(a.Command())
	s.WriteRune('(')
	for n := 0; n < a.ArgCount(); n++ {
		if n > 0 {
			s.WriteRune(',')
		}
		s.WriteRune('\'')
		s.WriteString(strings.Replace(a.Arg(n), "'", "\\'", -1))
		s.WriteRune('\'')
	}
	s.WriteRune(')')

	return s.String()
}

type bounds struct {
	l int
	r int
}

func emptyBounds() bounds {
	return bounds{
		l: -1,
		r: -1,
	}
}

func (b *bounds) IsValid() bool {
	return b.l >= 0 && b.r >= b.l
}
