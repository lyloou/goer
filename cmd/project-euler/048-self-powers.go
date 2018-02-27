package main

import (
	"math/big"
	"fmt"
)

func main() {

	var sum = big.NewInt(0)

	for i := 1; i <= 1000; i++ {
		var r = big.NewInt(1)
		for j := 1; j <= i; j++ {
			r = r.Mul(r, big.NewInt(int64(i)))
		}
		sum.Add(sum, r)
	}
	fmt.Println(sum)
}
