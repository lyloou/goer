package main

import (
	"fmt"
)

// https://gobyexample.com/closing-channels
func main() {
	jobs := make(chan string)
	done := make(chan bool) // add more than one channels, to indicate when jobs have done

	go func() {
		for {
			job, ok := <-jobs
			if ok {
				fmt.Println("do:" + job)
			} else {
				done <- true
				return
			}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("send jobs:", i)
			jobs <- fmt.Sprint(i)
		}
		close(jobs)
		fmt.Println("send all jobs")
	}()

	<-done
}
