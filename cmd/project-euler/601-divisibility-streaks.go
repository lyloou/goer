package main

import (
	"fmt"
	"math/big"

	"github.com/lyloou/goer/cmd/project-euler/common"
)

func main() {
	fmt.Println(common.Factorial(31))
	fmt.Println(P(3, 14))
	fmt.Println(P(6, 1000000))

	// TODO
	var i, sum int64
	for i = 1; i <= 31; i++ {
		power := getFourPower(i)
		count := P(i, power)
		fmt.Printf("P(%d, %d)=%d\n", i, power, count)
		sum += count
	}
	fmt.Println("result:", sum)
}

func getFourPower(length int64) int64 {
	// http://blog.csdn.net/u014609452/article/details/79115753
	// 4^x = 2^(2x) = (2^x)^2
	var c int64 = 1 << (uint64(length) << 1)
	return c
}

func getFourPower2(length int64) int64 {
	var r = big.NewInt(1)

	var i int64
	four := big.NewInt(4)
	for i = 1; i <= length; i++ {
		r = r.Mul(r, four)
	}
	return r.Int64()
}

func P(s, N int64) (count int64) {
	var n int64
	for n = 2; n < N; n++ {
		if !isNotStreak(n, s) {
			continue
		}
		if s == streak(n) {
			count ++
		}
	}
	return count
}

func streak(n int64) (count int64) {
	var i int64 = 1
	for n%i == 0 {
		count ++
		n ++
		i ++
	}
	return count
}

func isNotStreak(n, k int64) bool {
	return (n-1)%(k+1) != 0
}

func streak2(nn *big.Int) (count int64) {
	var one = big.NewInt(1)
	ii := big.NewInt(1)
	zz := big.NewInt(1)

	for zz.Mod(nn, ii).String() == "0" {
		count ++
		nn.Add(nn, one)
		ii.Add(ii, one)
	}
	return count
}
