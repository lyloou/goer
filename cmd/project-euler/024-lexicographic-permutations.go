package main

import "fmt"

var used []bool

func main() {
	m024of1()
}

// http://project-euler-answers-in-go.blogspot.com/2011/07/problem-24.html
// TODO: what's this?
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
