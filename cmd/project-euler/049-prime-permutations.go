package main

import (
	"fmt"
	"strconv"

	"github.com/lyloou/goer/cmd/project-euler/common"
)

func main() {

	var increment = 3330
	for i := 1001; i < 9999; i = i + 2 {
		if !common.IsPrime(int64(i)) {
			continue
		}

		j := i + increment
		if !common.IsPrime(int64(j)) {
			continue
		}

		if common.IsPermutation(strconv.Itoa(i), strconv.Itoa(j)) {
			k := j + increment
			if !common.IsPrime(int64(k)) {
				continue
			}

			if common.IsPermutation(strconv.Itoa(i), strconv.Itoa(k)) {
				fmt.Printf("%d%d%d\n", i, j, k)
			}
		}
	}

}
