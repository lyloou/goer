package main

import (
	"fmt"

	"github.com/lyloou/goer/cmd/project-euler/common"
)

// 有n个不同的factors

func main() {
	fmt.Println(firstConsecutiveInteger(1))
	fmt.Println(firstConsecutiveInteger(2))
	fmt.Println(firstConsecutiveInteger(3))
	fmt.Println(firstConsecutiveInteger(4)) // too slow, need to be optimized.
}

func firstConsecutiveInteger(length int) int {
	for num := 2; ; num++ {
		count := 0
		for i := 0; i < length; i++ {
			divisors := distinctPrimes(num, i)
			if len(divisors) != length {
				break
			}
			count ++
			if count == length {
				return num
			}
		}

	}
	return 0
}

func distinctPrimes(num int, i int) []int64 {
	divisors := common.GetDivisors(int64(num + i))
	for j := len(divisors) - 1; j >= 0; j-- {
		if !common.IsPrime(divisors[j]) {
			divisors = append(divisors[:j], divisors[j+1:]...)
		}
	}
	return divisors
}
