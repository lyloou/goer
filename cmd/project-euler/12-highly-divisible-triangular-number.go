package main

import (
	"fmt"
	"math"
)

func main() {
	var maxCount int64 = 0

	// 获得一个加法闭包
	f := Triangle()
	for {
		result := f()
		// 获得这个数的divisor个数
		count := getDivisorsCount(result)
		if count > maxCount {
			maxCount = count
			fmt.Println(result, count)

			if count >= 500 {
				break
			}
		}
	}
}

func getDivisors(division int64) (divisors [] int64) {
	var i int64
	for i = 1; i <= division/2; i++ {
		if division%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return
}

// 1: 1
// 3: 1,3
// 6: 1,2,3,6
// 10: 1,2,5,10
// 15: 1,3,5,15
// 21: 1,3,7,21
// 28: 1,2,4,7,14,28
// 思路：观察数的规律，可以先求出一半的divisors，然后返回个数的2倍
func getDivisorsCount(division int64) (count int64) {
	var i int64
	sqrt := math.Sqrt(float64(division))
	for i = 1; i <= int64(sqrt); i++ {
		if division%i == 0 {
			count++
		}
	}

	return count * 2
}

func Triangle() (func() int64) {
	var count, sum int64 = 0, 0
	return func() int64 {
		sum += count
		count++
		return sum
	}
}

/*
1 2
6 4
28 6
36 10
120 16
300 18
528 20
630 24
2016 36
3240 40
5460 48
25200 90
73920 112
157080 128
437580 144
749700 162
1385280 168
1493856 192
2031120 240
2162160 320
17907120 480
76576500 576
*/
