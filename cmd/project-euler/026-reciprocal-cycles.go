package main

import (
	"fmt"

	"github.com/lyloou/goer/pkg/util"
)

// https://www.mathblog.dk/project-euler-26-find-the-value-of-d-1000-for-which-1d-contains-the-longest-recurring-cycle/
func main() {
	maxLen := 0
	maxNum := 0
	maxArr := make([]int, 0)

	tmpArr := make([]int, 0)
	for num := 2; num < 1000; num++ {
		mod := 1

		for {
			mod = bigger(mod, num) % num
			if mod == 0 {
				tmpArr = tmpArr[:0]
				break
			}

			if util.IsContainsInIntSlice(tmpArr, mod) {
				if len(tmpArr) > maxLen {
					maxNum = num
					maxLen = len(tmpArr)
					maxArr = tmpArr
				}
				tmpArr = tmpArr[:0]
				break
			}

			tmpArr = append(tmpArr, mod)
		}
	}
	fmt.Println(maxNum, maxLen, maxArr)
}

func bigger(mod, num int) int {
	for {
		if mod < num {
			mod = mod * 10
		} else {
			return mod
		}
	}
}
