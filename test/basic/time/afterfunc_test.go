package time

import (
	"fmt"
	"testing"
	"time"
)

// https://yar999.gitbooks.io/gopl-zh/content/ch6/ch6-04.html
func TestAfterFunc(t *testing.T) {
	c := make(chan int)
	time.AfterFunc(3*time.Second, func() {
		t.Log("hi")
		t.Log(<-c)
	})
	time.AfterFunc(2*time.Second, IAmAFunc)

	c <- 1
}

func IAmAFunc() {
	fmt.Println("I am a function")
}
