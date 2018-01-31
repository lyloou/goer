package main

import (
	"fmt"
	"time"
)

func fibonacci(c chan int) {
	x, y := 1, 1
	for {
		c <- x
		x, y = y, x+y
		if x > 100 {
			break
		}
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}

func main() {
	c := make(chan int)
	go fibonacci(c)
	for i := range c {
		fmt.Println("-->", i)
	}

	tt := make(chan time.Time)
	timer := time.Timer{C: tt}
	fmt.Println("==>", timer)
	fmt.Println("==>", <-time.NewTimer(time.Millisecond * 5000).C)
}
