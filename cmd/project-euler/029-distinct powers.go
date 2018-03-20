package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(m029Of1(3, 4))
	fmt.Println(m029Of1(4, 3))

	m := make(map[string]string, 0)
	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			r := m029Of1(a, b)
			m[r] = r
		}
	}
	fmt.Println(len(m))
}

func m029Of1(a, b int) string {
	aa := big.NewInt(int64(a))
	r := big.NewInt(int64(a))
	for i := 1; i < b; i++ {
		aa = aa.Mul(aa, r)
	}
	return aa.String()
}
