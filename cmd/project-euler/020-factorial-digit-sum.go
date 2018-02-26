package main

import (
	"math/big"
	"fmt"
	"strconv"
)

func main() {
	f1()
	f2()
}

func f1() {
	var r = big.NewInt(1)
	for i := 1; i <= 100; i++ {
		r = r.Mul(r, big.NewInt(int64(i)))
	}

	sum := 0
	for _, v := range r.String() {
		c, err := strconv.Atoi(string(v))
		if err != nil {
			panic(err)
		}
		sum += c
	}
	fmt.Printf("int64 => %d\n", sum)
	fmt.Printf("int64 => %d\n", r)
}

// can not use float64, not exact
func f2() {
	var r float64 = 1
	for i := 1; i <= 100; i++ {
		r = r * float64(i)
	}
	fmt.Printf("float => %f\n", r)
}
