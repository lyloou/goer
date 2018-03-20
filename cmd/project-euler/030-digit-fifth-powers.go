package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	var a int
	for a = 2; a < 10000000; a++ {
		if m030Of1(a) {
			sum += a
			fmt.Println(a, sum)
		}
	}
	fmt.Println(sum)
}

func m030Of1(a int) bool {
	aa := strings.Split(fmt.Sprintf("%v", a), "")
	r := 0
	for _, v := range aa {
		vv, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		r += vv * vv * vv * vv * vv
	}
	return r == a
}
