package main

import (
	"fmt"

	"github.com/lyloou/goer/cmd/project-euler/common"
)

// http://blog.csdn.net/u014609452/article/details/79115753
func main() {
	var i, result int64
	var lcm, nxt, tmp int64 = 1, 0, 0
	for i = 1; i <= 31; i++ {
		n := common.FourPower(i)
		nxt = lcm * (i + 1) / common.Gcd(lcm, i+1)
		tmp = Cnt(n-2, lcm) - Cnt(n-2, nxt)
		fmt.Printf("%v, %v, %v\n", i, n, tmp)
		result += tmp
		fmt.Println("====================>", lcm)
		lcm = nxt
	}
	fmt.Println(result)
}

func Cnt(a, b int64) int64 {
	return a / b
}
