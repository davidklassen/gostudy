package memory

import (
	"fmt"
)

type foo int

type bar interface {
	val() int
}

func (f foo) val() int {
	return int(f)
}

func AssignIface() {
	f := foo(4)
	var b bar = f
	f = foo(5)
	fmt.Println(b.val())

	var fn func()

	fn()
}
