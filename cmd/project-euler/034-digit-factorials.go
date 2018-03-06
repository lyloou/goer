package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lyloou/goer/cmd/project-euler/common"
)

func main() {
	var i, result int = 3, 0
	for ; i < 1000000; i++ {
		str := strings.Split(strconv.Itoa(i), "")
		sum := 0
		for _, v := range str {
			vv, _ := strconv.Atoi(v)
			sum += int(common.Factorial(int64(vv)))
		}

		if sum == i {
			fmt.Println("this ok", i)
			result += sum
		}
		fmt.Println(result)
	}
}
