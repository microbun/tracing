package stackerr

import (
	"fmt"
	"io"
	"runtime"
)

type frame uintptr

func (f frame) pc() uintptr {
	return uintptr(f)
}

func (f frame) Format(s fmt.State, verb rune) {
	pc := f.pc()
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		file, line := fn.FileLine(pc)
		fmt.Fprintf(s, "%s\n\t%s:%d\n", fn.Name(), file, line)
	} else {
		fmt.Fprintf(s, "unkown stack\n")
	}
}

//type stack uintptr
type errorStack struct {
	message string
	cause   error
	stack   []frame
}

func (e *errorStack) Error() string {
	if e.cause != nil && e.message == "" {
		return e.cause.Error()
	}
	return e.message
}

func (e *errorStack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		// fmt.Fprintf(s, "%+v\n", e.err)
		fmt.Fprintf(s, "%s\n", e.Error())

		for _, f := range e.stack {
			fn := runtime.FuncForPC(f.pc())
			if fn != nil {

				fmt.Fprintf(s, "%v", f)
			} else {
				fmt.Fprintf(s, "unkown stack\n")
			}
		}
		if v, ok := e.cause.(*errorStack); ok {
			fmt.Fprintf(s, "cause by: %v", v)
		}
	case 's':
		io.WriteString(s, e.Error())
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	}
}

//New returns an error with stack
func New(message string) error {

	return &errorStack{
		message: message,
		cause:   nil,
		stack:   callers(0),
	}
}

//RawError return the raw error
func RawError(err error) error {
	if err == nil {
		return err
	}
	for {
		if e, ok := err.(*errorStack); ok {
			err = e.cause
		} else {
			return err
		}
	}

}

//Cause return a cuase
func Cause(err error) error {
	if err == nil {
		return err
	}
	if e, ok := err.(*errorStack); ok {
		return e.cause
	}
	return err
}

//WithStack return a err with a stack trace
func WithStack(err error) error {
	if err == nil {
		return err
	}
	if e, ok := err.(*errorStack); ok {
		pc, _, _, _ := runtime.Caller(2)
		e.stack = append(e.stack, frame(pc))
		return e
	}

	return &errorStack{
		cause: err,
		stack: callers(0),
	}
}

func WithMessage(err error, message string) error {
	return &errorStack{
		message: message,
		cause:   err,
		stack:   callers(0),
	}
}

func callers(n int) []frame {
	n = n + 3
	pc2, _, _, _ := runtime.Caller(n)
	pc1, _, _, _ := runtime.Caller(n - 1)
	pcs := make([]frame, 0)
	pcs = append(pcs, frame(pc1))
	pcs = append(pcs, frame(pc2))
	return pcs
}
