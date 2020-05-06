package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	pow := decimal.New(2, 0).Pow(decimal.New(7830457, 0)).Mul(decimal.New(28433, 0)).Add(decimal.New(1, 0))
	result := pow.String()
	length := len(result)
	fmt.Println(result[length-10 : length])
}
