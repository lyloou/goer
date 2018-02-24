package main

import "fmt"

func main() {
	var i int64
	max := 0
	for i = 1; i < 1000000; i++ {
		count := chain(i)
		if count > max {
			max = count
			fmt.Printf("num:%6d, count:%3d\n", i, max)
		}
	}
}

func chain(num int64) (int) {
	tmp := num
	for i := 0; ; i++ {
		if isEven(tmp) {
			tmp = doEven(tmp)
		} else {
			if tmp == 1 {
				return i
			}
			tmp = doOdd(tmp)
		}
	}
}

func isEven(num int64) bool {
	return num%2 == 0
}

func doEven(num int64) int64 {
	return num / 2
}

func doOdd(num int64) int64 {
	return 3*num + 1
}
