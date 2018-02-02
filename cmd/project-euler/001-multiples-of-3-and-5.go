package main

import (
	"fmt"
)

// https://projecteuler.net/problem=1
func main() {
	max := 1000
	sum := 0
	for i := 3; i < max; i++ {
		if i%3 == 0 || i%5 == 0 {
			fmt.Println(i)
			sum += i
		}
	}
	fmt.Println("sum:", sum)
}
