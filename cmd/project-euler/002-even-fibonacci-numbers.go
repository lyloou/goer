package main

import "fmt"

// https://projecteuler.net/problem=2

func main() {
	var maxData, a, b, sum int64
	maxData = 4000000
	a = 1
	b = 2
	fibNum := make([]int64, 0)
	fibNum = append(fibNum, a)
	for b < maxData {
		fibNum = append(fibNum, b)
		a, b = b, a+b
	}

	for _, v := range fibNum {
		if v%2 == 0 {
			sum += v
		}
	}

	fmt.Println(fibNum)
	fmt.Println(sum)
	evenFibonacciNumbers()
}

func evenFibonacciNumbers() {
	var maxData, a, b, sum int64
	maxData = 4000000
	a = 1
	b = 2
	for b < maxData {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, a+b
	}
	fmt.Println(sum)
}
