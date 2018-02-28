package basic

import (
	"container/list"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	list := list.New()
	for i := 0; i < 10; i++ {
		list.PushBack(i)
	}
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println(list)
}
