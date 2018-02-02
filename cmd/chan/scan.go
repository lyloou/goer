// https://www.golang-book.com/books/intro/10

package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)

	fmt.Println(cap(c))
	t := time.After(time.Millisecond * 2000)
	fmt.Println(<-t)
	fmt.Println(cap(t))

	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}

func printer(c <-chan string) {
	for v := range c {
		time.Sleep(time.Millisecond * 500)
		fmt.Println(v)
	}
}

func ponger(c chan<- string) {
	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("pong:%d", i)
	}
	// close(c) // can't send data to a closed chan

}

func pinger(c chan<- string) {
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("ping:%d", i)
	}
}
