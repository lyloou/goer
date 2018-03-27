package main

import (
	"fmt"
	"strconv"

	"github.com/lyloou/goer/cmd/project-euler/common"
)

func main() {
	max := 0
	for i := 1; i < 10000; i++ {
		tmp := ""
		for j := 1; j < 10 && len(tmp) < 9; j++ {
			tmp += strconv.Itoa(i * j)
		}

		if common.IsPandigital2(tmp) {
			if v, ok := strconv.Atoi(tmp); ok == nil {
				if max < v {
					max = v
				}
				fmt.Println(max, i)
			}
		}
	}
	fmt.Println(max)
}
