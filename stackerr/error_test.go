package stackerr

import (
	"testing"
)

func TestB(t *testing.T) {
	err := d()

	// PrintStack(Cause(err))
	// fmt.Println("============")
	PrintStack(err)
	// fmt.Println("============")
	// debug.PrintStack()
}

func a() error {
	// _, err := os.Open("1")
	// if err != nil {
	// 	return WithStack(err)
	// }
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
