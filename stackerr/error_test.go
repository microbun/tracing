package stackerr

import (
	"fmt"
	"testing"
)

func TestB(t *testing.T) {
	err := d()

	fmt.Printf("%v", err)
}

func a() error {
	return New("xxx")
}

func b() error {
	err := a()
	return WithStack(err)
}

func c() error {
	err := b()
	return WithStack(err)
}

func d() error {
	err := c()
	return WithStack(err)
}
