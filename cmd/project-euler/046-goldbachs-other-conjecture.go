package main

import (
	"fmt"
	"math"

	"github.com/lyloou/goer/cmd/project-euler/common"
)

func main() {
	var i, j int64
	for i = 9; ; i += 2 {
		if common.IsPrime(i) {
			continue
		}

	inner:
		for j = i - 2; j > 0; j -= 2 {
			if common.IsPrime(j) {
				m := (i - j) / 2
				mm := math.Sqrt(float64(m))
				if int64(mm)*int64(mm) == m {
					fmt.Println(i, j)
					break inner
				}
			}
		}

		if j <= 0 {
			fmt.Println("got it:", i)
			return
		}
	}
}

//9 = 7 + 2×12
//15 = 7 + 2×22
//21 = 3 + 2×32
//25 = 7 + 2×32
//27 = 19 + 2×22
//33 = 31 + 2×12
