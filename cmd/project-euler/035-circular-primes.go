package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lyloou/goer/cmd/project-euler/common"
)

func main() {
	var c, i int64 = 1, 0 // i start with 3.  2 is prime, so c=1
	for i = 3; i < 1000000; i++ {
		if common.IsPrime(i) {
			if isCircularPrimes(i) {
				c++
			}
		}
	}

	fmt.Println("==>", c)
}

func isCircularPrimes(num int64) bool {
	a := strconv.Itoa(int(num))
	if maybeEven(a) {
		return false
	}

	arr := strings.Split(a, "")
	for i := 1; i < len(arr); i++ {
		arr = append(arr[1:], arr[:1]...)
		n, _ := strconv.Atoi(strings.Join(arr, ""))
		if !common.IsPrime(int64(n)) {
			return false
		}
	}
	fmt.Println(num)
	return true
}

func maybeEven(s string) bool {
	return strings.Contains(s, "0") ||
		strings.Contains(s, "2") ||
		strings.Contains(s, "4") ||
		strings.Contains(s, "6") ||
		strings.Contains(s, "8")
}
