package main

import "fmt"

// 1
// 3	5	7	9
// 13	17	21	25
// 31	37	43	49
func main() {

	x := 1
	z := 2
	sum := x
	for z <= 1001 {
		for i := 0; i < 4; i++ {
			x += z
			sum += x
			fmt.Println(x)
		}
		z = z + 2
	}
	fmt.Println(sum)
}
