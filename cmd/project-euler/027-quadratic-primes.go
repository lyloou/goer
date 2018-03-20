package main

import (
	"fmt"
	"math"

	"github.com/lyloou/goer/cmd/project-euler/common"
)

func main() {
	fmt.Println(isQuadraticPrimes(1, 41))
	fmt.Println(isQuadraticPrimes(-79, 1601))

	var i, maxResult, vr, vvr int64
	primes := make([]int64, 0)
	for i = 1; i < 1000; i++ {
		if common.IsPrime(i) {
			primes = append(primes, i, -i)
		}
	}
	fmt.Println(len(primes), primes)

	for i := 0; i < len(primes); i++ {
		for j := 0; j < len(primes); j++ {
			v, vv := primes[i], primes[j]
			if isQuadraticPrimes(v, vv) {
				result := int64(math.Abs(float64(v * vv)))
				if maxResult < result {
					maxResult = result
					vr = v
					vvr = vv
				}
			}
		}
	}
	fmt.Println(vr*vvr, vr, vvr)
}

func isQuadraticPrimes(a, b int64) bool {
	var n int64
	aa := int64(math.Abs(float64(a)))
	for n = 0; n < aa; n++ {
		if !common.IsPrime(n*n + a*n + b) {
			return false
		}
	}
	return true
}
