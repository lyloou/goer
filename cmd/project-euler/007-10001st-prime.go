package main

import (
	"fmt"
	"github.com/lyloou/goer/cmd/project-euler/common"
)

func main() {
	//t := 6
	t := 10001
	c := 0
	for i := 2; i < 1000000; i++ {
		if common.IsPrime(int64(i)) {
			c ++
		}
		if c >= t {
			fmt.Println(i)
			break
		}
	}

}
