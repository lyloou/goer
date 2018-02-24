package main

import (
	"math/big"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	a := big.NewInt(1)

	for i := 0; i < 1000; i++ {
		a.Add(a, a)
	}
	fmt.Println(a.String())

	sum := 0
	datas := strings.Split(a.String(), "")
	for _, data := range datas {
		v, err := strconv.Atoi(data)
		if err != nil {
			panic(err)
		}
		sum += v
	}
	fmt.Println(sum)
}
