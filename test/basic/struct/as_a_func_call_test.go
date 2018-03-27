package _struct

import "testing"

// https://yar999.gitbooks.io/gopl-zh/content/ch6/ch6-04.html
func TestAsAFuncAll(t *testing.T) {
	var b B
	b.SayHello()

	f := (*B).SayHello
	f(nil)

	f2 := (*B).SayHelloWithB
	f2(nil)

	f3 := (*B).SayHelloWithBCall
	f3(nil) // panic: runtime error: invalid memory address or nil pointer dereference [recovered]
}
