package main

import (

"fmt"

"github.com/lyloou/goer/cmd/project-euler/common"

)


func main() {
	fmt.Println( )
	return
	var i int64
//for1:
	for i = 9; ; i += 2 {
		if common.IsPrime(i) {
			continue
		}

	for2:
		for j := i - 2; j > 0; j -= 2 {
			if common.IsPrime(j) {
				if (j-i)%2 != 0 {
					panic(fmt.Sprintf("1. error  i:%v, j:%v", i, j))
					return
				}

				fmt.Println(i, j)
				m := (j - i) / 2

			for3:
				for k := 1; ; k ++ {
					if m == int64(k*k) {
						fmt.Printf("%v\t= %v + 2*(%v*%v)\n", i, j, k, k)
						break for2
					}
					if m > int64(k*k) {
						break for3
					}
				}
			}
		}

	}

}

//9 = 7 + 2×12
//15 = 7 + 2×22
//21 = 3 + 2×32
//25 = 7 + 2×32
//27 = 19 + 2×22
//33 = 31 + 2×12
