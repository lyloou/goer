package main

import (
	"fmt"

	"github.com/lyloou/goer/cmd/project-euler/common"
)

// https://blog.csdn.net/qunxingvip/article/details/46352005
func main() {


}

/*

https://stackoverflow.com/questions/25696076/consecutive-prime-sum
A smaller example: if you were finding the largest sum of consecutive primes with sum less than 10, then that is 3 + 5 = 8, not 2 + 3 = 5.
It might not (and is not) the case that you always get the largest sum by adding all the primes starting at 2.

*/
func errorWay1() {
	var max int64 = 1000000
	var i, sum int64 = 0, 0
	for i = 1; ; i = i + 1 {
		if sum > max {
			break
		}
		if common.IsPrime(i) {
			sum += i
			if common.IsPrime(sum) {
				fmt.Println(i, sum)
			}
		}
	}
	fmt.Println(common.IsPrime(999983))
}
