package main

import (
	"github.com/lyloou/goer/cmd/project-euler/common"
	"fmt"
)

const maxSum = 2000000

func main() {
	sum := 0
	for i := 2; i < maxSum; i++ {
		if common.IsPrime(int64(i)) {
			sum += i
		}
	}

	fmt.Println("sum:", sum)
}
