package main

import (
	"fmt"
	"math"
)

func main() {
	//t := 6
	t := 10001
	c := 0
	for i := 2; i < 1000000; i++ {
		if isPrime(i) {
			c ++
		}
		if c >= t {
			fmt.Println(i)
			break
		}
	}

}

func isPrime(n int) bool {
	s := int(math.Sqrt(float64(n))) + 1
	for i := int(2); i < s; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
