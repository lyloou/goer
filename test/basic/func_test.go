package basic

import (
	"testing"
	"fmt"
)

type func1 func(i int)

// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/03.4.md
// 类似HandlerFunc，测试func的强制转换功能
func (f func1) call(i int) {
	fmt.Printf("\n==>%d<==\n", i)
	f(i)
	fmt.Printf("==>%d<==\n", i)
}

func TestFunc1(t *testing.T) {
	ft(1)
	f := func1(ft)
	f(2)
	f.call(3)
}

func ft(i int) {
	fmt.Println("from func t: ", i)
}
