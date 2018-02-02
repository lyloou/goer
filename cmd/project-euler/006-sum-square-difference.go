package main

import "fmt"

func main() {
	var max, tmp1, tmp2, result = 100, 0, 0, 0
	for i := 0; i <= max; i++ {
		tmp1 += i * i
	}

	for i := 0; i <= max; i++ {
		tmp2 += i
	}
	tmp2 = tmp2 * tmp2

	result = tmp2 - tmp1
	fmt.Println(tmp1, tmp2, result)
}
