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

func open() (*os.File, error) {
	f, err := os.Open("a.txt")
	if err != nil {
		return nil, stackerr.WithStack(err)
	}
	return f, nil
}

func main() {

	f, err := open()
	if err != nil {
		stackerr.PrintStack(err)
		return
	}
	fmt.Println(f.Name())

}

```

Result:

```base
error: open a.txt: no such file or directory
main.open
        /home/micro/src/main.go:13
main.main
        /home/micro/src/main.go:20

```
