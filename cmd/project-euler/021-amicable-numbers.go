package main

import (
	"github.com/lyloou/goer/cmd/project-euler/common"
	"fmt"
)

func main() {
	var i, sum int64
	for i = 1; i <= 10000; i++ {
		if isAmicable(i) {
			fmt.Println(i)
			sum += i
		}
	}
	fmt.Println(sum)
}

func sumDivisors(divisors []int64) (sum int64) {
	for _, v := range divisors {
		sum += v
	}
	return
}

func isAmicable(num int64) bool {
	sum := sumDivisors(common.GetDivisors(num))
	if sum == num {
		return false
	}
	return sumDivisors(common.GetDivisors(sum)) == num
}
