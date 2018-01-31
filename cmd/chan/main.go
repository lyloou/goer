package main

import (
	"time"
	"fmt"
)

// test single channel
// http://www.cnblogs.com/baiyuxiong/p/4545028.html
// https://yuerblog.cc/2017/07/15/golang-single-direction-channel/
// http://legendtkl.com/2017/07/30/understanding-golang-channel/
func main() {
	c := make(chan int)
	go send(c)
	go recv(c)
	time.Sleep(3 * time.Second)
}
func recv(c <-chan int) {
	for v := range c {
		time.Sleep(299 * time.Millisecond)
		fmt.Println(v)
	}
}
func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
}
