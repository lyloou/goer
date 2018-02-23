package main

import (
	"fmt"
	"math/big"
	"bufio"
	"github.com/lyloou/goer/pkg/util"
	"github.com/lyloou/goer/cmd/project-euler/handler"
)

// https://projecteuler.net/problem=13
// Large sum
// Problem 13
// Work out the first ten digits of the sum of the following one-hundred 50-digit numbers.

// link: https://gist.github.com/kendellfab/7417164
// https://stackoverflow.com/questions/10747411/using-big-integer-values-in-go-parseint-only-converts-up-to-2147483647

// $ cd project-euler
// $ go run 013-largesum.go

func main() {
	err := handler.HandleScannerWithFilename("013-largesum.data", func(scanner *bufio.Scanner) error {
		a := big.NewInt(0)
		b := big.NewInt(0)
		for scanner.Scan() {
			a.SetString(scanner.Text(), 10)
			b.Add(a, b)
		}
		fmt.Println(a)
		fmt.Println(util.Substr(b.String(), 0, 10))
		return nil
	})

	if err != nil {
		panic(err)
	}
}
