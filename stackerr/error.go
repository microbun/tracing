package stackerr

import (
	"errors"
	"fmt"
	"io"
	"runtime"
)

//type stack uintptr
type errorStack struct {
	err   error
	stack []uintptr
}

func (e *errorStack) Error() string {
	return e.err.Error()
}

func (e *errorStack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		// fmt.Fprintf(s, "%+v\n", e.err)
		fmt.Fprintf(s, "error:%s\n", e.Error())
		for _, pc := range e.stack {
			fn := runtime.FuncForPC(pc)
			if fn != nil {
				file, line := fn.FileLine(pc)
				fmt.Fprintf(s, "%s\n\t%s:%d\n", fn.Name(), file, line)
			} else {
				fmt.Fprintf(s, "unkown stack\n")
			}
		}
		// fallthrough
	case 's':
		io.WriteString(s, e.Error())
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	}
}

//New returns an error with stack
func New(message string) error {

	return &errorStack{
		err:   errors.New(message),
		stack: callers(0),
	}
}

//RawError return the raw error
func RawError(err error) error {
	if err == nil {
		return err
	}
	if e, ok := err.(*errorStack); ok {
		return e.err
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
		e.stack = append(e.stack, pc)
		return e
	}
	// caller := [8]uintptr{}

	return &errorStack{
		err:   err,
		stack: callers(0),
	}
}

func callers(n int) []uintptr {
	n = n + 3
	pc2, _, _, _ := runtime.Caller(n)
	pc1, _, _, _ := runtime.Caller(n - 1)
	pcs := make([]uintptr, 0)
	pcs = append(pcs, pc1)
	pcs = append(pcs, pc2)
	return pcs
}

// //PrintStack print a stack trace
// func PrintStack(err error) {
// 	fmt.Printf("error: %s\n", err.Error())
// 	if e, ok := err.(*errorStack); ok {
// 		for _, pc := range e.stack {
// 			printPC(pc)
// 		}
// 	}
// }

// func printPC(pc uintptr) {
// 	fn := runtime.FuncForPC(pc)
// 	if fn != nil {
// 		file, line := fn.FileLine(pc)
// 		fmt.Printf("%s\n\t%s:%d\n", fn.Name(), file, line)
// 	} else {
// 		fmt.Printf("unkown stack\n")
// 	}
// }
