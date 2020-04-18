# tracing

Simple error handling and tracing for Golang

Example:

```golang
package main

import (
	"fmt"
	"os"

	"github.com/microbun/tracing/stackerr"
)

func funca() error {
	_, err := os.Open("test.txt")
	if err != nil {
		return stackerr.WithStack(err)
	}
	return nil
}

func funcb() error {
	err := funca()
	return stackerr.WithStack(err)
}

func funcc() error {
	err := funcb()
	return stackerr.WithStack(err)
}

func funcd() error {
	return stackerr.New("is error")
}

func main() {
	err := funcc()
	if err != nil {

		//get raw error
		rawErr := stackerr.RawError(err)

		if os.IsNotExist(rawErr) {
			//print raw error message
			fmt.Printf("funcc() raw error with message: \n%v\n\n\n", rawErr)

			//print error with stack
			fmt.Printf("funcc() error with stack: \n%v", err)
		}
	}
	fmt.Println("\n\n")
	err1 := funcd()
	if err1 != nil {
		fmt.Printf("funcd() error with stack: \n%v", err1)
	}

}

```

Console Output:

```base
funcc() raw error with message: 
open test.txt: no such file or directory


funcc() error with stack: 
error:open test.txt: no such file or directory
main.funca
	/home/micro/go/src/main.go:13
main.funcb
	/home/micro/go/src/main.go:19
main.funcc
	/home/micro/go/src/main.go:24
main.main
	/home/micro/go/src/main.go:33



funcd() error with stack: 
error:is error
main.funcd
	/home/micro/go/src/main.go:29
main.main
	/home/micro/go/src/main.go:48


```
