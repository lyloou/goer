package main

import (
	"fmt"
	"math/big"
	"os"
	"bufio"
	"github.com/lyloou/goer/pkg/util"
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
	f, err := os.Open("013-largesum.data")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	a := big.NewInt(0)
	b := big.NewInt(0)
	for scanner.Scan() {
		a.SetString(scanner.Text(), 10)
		b.Add(a, b)
	}
	fmt.Println(a)
	fmt.Println(util.Substr(b.String(), 0, 10))

}
