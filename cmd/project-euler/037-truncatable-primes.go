package main

import (
	"fmt"

	"github.com/lyloou/goer/cmd/project-euler/common"
)

// The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right, and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.
// Find the sum of the only eleven primes that are both truncatable from left to right and right to left.
// NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.

// 读懂题目的意思很重要，不然无法下手啊！
// https://zh.wikipedia.org/wiki/%E5%8F%AF%E6%88%AA%E7%9F%AD%E8%B3%AA%E6%95%B8
// https://blog.csdn.net/LuoQuHen/article/details/45200539
// 找到这11个数字（既是可左截短质数，又是可右截短质数），并求和。

const MAX_NUM = 10000000000000

func main() {
	fmt.Println(isTruncatablePrime(3797))

	var i, sum int64
	for i = 10; i < MAX_NUM; i++ {
		if isTruncatablePrime(i) {
			sum += i
			fmt.Printf("i=%d, sum=%d\n", i, sum)
		}
	}
}

// 分析：
// left to right: 3797, 797, 97, and 7.
// 7    = 3797 % 10
// 97   = 3797 % 100
// 797  = 3797 % 1000

// right to left: 3797, 379, 37, 3
// 3    = 3797 / 1000
// 37   = 3797 / 100
// 379  = 3797 / 10
func isTruncatablePrime(num int64) bool {
	if !common.IsPrime(num) {
		return false
	}
	var i, r int64
	for i = 10; i < MAX_NUM; i = i * 10 {
		r = num / i  // 二合一，1
		m := num % i // 二合一，2
		if r <= 0 {
			break
		}
		if !common.IsPrime(r) || !common.IsPrime(m) {
			return false
		}
	}
	return true
}

func isTruncatablePrime2(num int64) bool {
	if !common.IsPrime(num) {
		return false
	}
	var i, r int64
	for i = 10; i < MAX_NUM; i = i * 10 {
		r = num / i // 二合一，1
		if r <= 0 {
			break
		}
		if !common.IsPrime(r) {
			return false
		}
	}

	for i = 10; i < MAX_NUM; i = i * 10 {
		m := num % i // 二合一，2
		if !common.IsPrime(m) {
			return false
		}
	}
	return true
}
