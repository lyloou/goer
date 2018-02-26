package main

import (
	"fmt"
	"math/big"
)

// https://projecteuler.net/problem=25

func main() {
	var a, b *big.Int
	a = big.NewInt(1)
	b = big.NewInt(1)
	c := 1
	for {
		c++
		a, b = b, a.Add(a, b)
		if len(a.String()) == 1000 {
			break
		}
	}
	fmt.Println(c)
}
