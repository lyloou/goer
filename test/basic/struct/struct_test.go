package _struct

import (
	"fmt"
	"testing"
)

// https://yar999.gitbooks.io/gopl-zh/content/ch6/ch6-03.html
func TestStruct(t *testing.T) {
	var c C
	//c.Add(23) // ambiguous reference
	c.A.Add(1)
	fmt.Printf("%#v\n", c)

	c.Minus(32)
	fmt.Printf("%#v\n", c)

	c.SayHello()
}

type A struct {
	X int
}

type B struct {
	X int
}

type C struct {
	A
	B
}

func (a *A) Add(b int) {
	fmt.Println("from A")
	a.X = a.X + b
}

func (b *B) Add(a int) {
	fmt.Println("from B")
	b.X = b.X + a
}

func (b *B) Minus(a int) {
	fmt.Println("from B")
	b.X = b.X - a
}

func (b *B) SayHello() {
	fmt.Println("Hello")
}

func (b *B) SayHelloWithB() {
	fmt.Println("Hello:", b)
}

func (b *B) SayHelloWithBCall() {
	fmt.Println("Hello:", b.X)

}
