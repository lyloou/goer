package main

import (
	"fmt"

	"github.com/lyloou/goer/cmd/project-euler/common"
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

func isAmicable(num int64) bool {
	sum := common.SumDivisors(common.GetDivisors(num))
	if sum == num {
		return false
	}
	return common.SumDivisors(common.GetDivisors(sum)) == num
}
