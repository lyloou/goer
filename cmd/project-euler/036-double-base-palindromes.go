package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isDoubleBasePalindromes(2))
	fmt.Println(isDoubleBasePalindromes(7))
	fmt.Println(isDoubleBasePalindromes(585))

	sum := 0
	for i := 0; i < 1000000; i++ {
		if isDoubleBasePalindromes(i) {
			fmt.Println(i)
			sum += i
		}
	}

	fmt.Println("sum:", sum)
}

func isDoubleBasePalindromes(num int) bool {
	return isPalindromesStr(strings.Split(fmt.Sprintf("%d", num), "")) &&
		isPalindromesStr(strings.Split(fmt.Sprintf("%b", num), ""))
}

func isPalindromesStr(strs []string) bool {
	l := len(strs)
	if l < 2 {
		return true
	}

	left := l/2 - 1
	right := l / 2
	if l%2 == 1 { // if length is odd, right move one step
		right = right + 1
	}
	for i := 0; i < l/2; i++ {
		if strs[left-i] != strs[right+i] {
			return false
		}
	}

	return true
}
