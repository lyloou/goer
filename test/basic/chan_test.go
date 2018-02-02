package basic

import (
	"testing"
	"fmt"
	"time"
)

func TestChan(t *testing.T) {
	stop := make(chan bool)
	go func(c chan bool) {
		time.Sleep(time.Second * 3)
		fmt.Println("1")
		go func() {
			time.Sleep(time.Second * 3)
			fmt.Println("2")
			close(c)
		}()
	}(stop)

	// 当所有的goroutine都结束了，却还没有执行到这里来，就抛出异常出来。
	// fatal error: all goroutines are asleep - deadlock!
	v, ok := <-stop
	fmt.Println(v, ok)

	if ok {
		t.Error("Channel is not closed")
	}

}
