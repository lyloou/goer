package main

import (
	"fmt"
	"time"
)

// link: github.com/chanxuehong/util/math.MaxInt
const (
	IntSize  = 32 << (^uint(0) >> 63)
	UintSize = 32 << (^uint(0) >> 63)
)

const (
	MaxInt  = 1<<(IntSize-1) - 1
	MinInt  = -1 << (IntSize - 1)
	MaxUint = 1<<UintSize - 1
)

func main() {
	fmt.Println(time.Now())
	for i := 1; i < MaxInt; i++ {
		if isSmallest(i) {
			fmt.Println("find it:", i)
			break
		}
	}
	fmt.Println(time.Now())
}

func isSmallest(num int) bool {
	for i := 2; i <= 20; i++ {
		if num%i != 0 {
			return false
		}

	}
	return true
}
