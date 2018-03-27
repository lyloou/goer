package basic

import (
	"fmt"
	"testing"
)

func TestStack_PushPop(t *testing.T) {

	tests := []struct {
		input string
		want  string
	}{
		{"a", "a"},
		{"b", "b"},
		{"c", "c"},
	}

	var stack Stack
	for _, v := range tests {
		stack.Push(v.input)
		fmt.Println("==>", stack)

		got, err := stack.Pop()
		if err != nil && got != v.want {
			t.Errorf("stack.Pop() = %v, want:%v", got, v.want)
		}
	}

	got, err := stack.Pop()
	fmt.Println(got, err == EmptyStackError)
}

