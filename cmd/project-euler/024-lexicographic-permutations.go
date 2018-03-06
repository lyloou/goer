package main

import (
	"fmt"

	"github.com/lyloou/goer/cmd/project-euler/common"
)

func main() {
	m024of5()
}

// refer to func m024of4
// link: https://en.wikipedia.org/wiki/Factorial_number_system#Permutations
func m024of5() {
	n := 1000000 - 1
	factorial := 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
	factoid := make([]int, 10)
	for i := 9; i > 0; i-- {
		factoid[i] = n / factorial
		n %= factorial
		factorial /= i
	}
	fmt.Println(factoid)
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 9; i >= 0; i-- {
		a := factoid[i]
		fmt.Print(arr[a])
		arr = append(arr[:a], arr[a+1:]...)
	}
}

func m024of4() {
	// [0 1 1 2 1 5 2 6 6 2]
	result :=
		2*common.Factorial(9) +
			6*common.Factorial(8) +
			6*common.Factorial(7) +
			2*common.Factorial(6) +
			5*common.Factorial(5) +
			1*common.Factorial(4) +
			2*common.Factorial(3) +
			1*common.Factorial(2) +
			1*common.Factorial(1) +
			0*common.Factorial(0)
	fmt.Println(result)
}

// refer to func m024of2
// link: https://en.wikipedia.org/wiki/Factorial_number_system
func m024of3() {
	n := 1000000 - 1
	factorial := 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
	factoid := make([]int, 10)
	for i := 9; i > 0; i-- {
		factoid[i] = n / factorial
		n %= factorial
		factorial /= i
	}
	fmt.Println(factoid)

	for i := 1; i < 10; i++ {
		for j := i - 1; j >= 0; j-- {
			if factoid[j] >= factoid[i] {
				factoid[j] ++
			}
		}
	}

	for i := 9; i >= 0; i-- {
		fmt.Print(factoid[i])
	}

}

var used []bool
// http://project-euler-answers-in-go.blogspot.com/2011/07/problem-24.html
// Q: what's mean?  A: solve it by factorial
// https://www.mathblog.dk/project-euler-24-millionth-lexicographic-permutation/
func m024of2() {
	used = make([]bool, 10)
	for i := range used {
		used[i] = false
	}
	remaining := 999999
	perm := 1
	index := 0
	for i := 9; i > 0; i-- {
		perm = 1
		for j := i; j > 0; j-- {
			perm *= j
		}
		index = remaining / perm
		remaining = remaining % perm
		j := 0
		for j = 0; j < index || used[j]; j++ {
			if used[j] {
				index++
			}
		}
		used[j] = true
		print(j)
	}
	println(remaining)
}

// 按照组合的方式来思考：10*9*8*7*6*5*4*3*2*1
func m024of1() {
	var arr = make([]int, 0)
	const C = 10
	const LAST = 1000000
	var count int64 = 0

out:
	for a0 := 0; a0 < C; a0++ {
		arr = arr[:0]
		if contains(a0, arr) {
			continue
		}
		arr = append(arr, a0)

		for a1 := 0; a1 < C; a1++ {
			arr = arr[:1]
			if contains(a1, arr) {
				continue
			}
			arr = append(arr, a1)

			for a2 := 0; a2 < C; a2++ {
				arr = arr[:2]
				if contains(a2, arr) {
					continue
				}
				arr = append(arr, a2)

				for a3 := 0; a3 < C; a3++ {
					arr = arr[:3]
					if contains(a3, arr) {
						continue
					}
					arr = append(arr, a3)

					for a4 := 0; a4 < C; a4++ {
						arr = arr[:4]
						if contains(a4, arr) {
							continue
						}
						arr = append(arr, a4)

						for a5 := 0; a5 < C; a5++ {
							arr = arr[:5]
							if contains(a5, arr) {
								continue
							}
							arr = append(arr, a5)

							for a6 := 0; a6 < C; a6++ {
								arr = arr[:6]
								if contains(a6, arr) {
									continue
								}
								arr = append(arr, a6)

								for a7 := 0; a7 < C; a7++ {
									arr = arr[:7]
									if contains(a7, arr) {
										continue
									}
									arr = append(arr, a7)

									for a8 := 0; a8 < C; a8++ {
										arr = arr[:8]
										if contains(a8, arr) {
											continue
										}
										arr = append(arr, a8)

										for a9 := 0; a9 < C; a9++ {
											arr = arr[:9]
											if contains(a9, arr) {
												continue
											}
											arr = append(arr, a9)

											count ++
											if count == LAST {
												fmt.Printf("%v%v%v%v%v%v%v%v%v%v", a0, a1, a2, a3, a4, a5, a6, a7, a8, a9)
												break out
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func contains(i int, arr []int) bool {
	for _, v := range arr {
		if v == i {
			return true
		}
	}
	return false
}
