package main

import (
	"strings"
	"strconv"
	"fmt"
	"time"
)

func main() {
	var top, bottom, tmp, max = 999, 100, 0, 0

	fmt.Println(time.Now())
	for a := top; a >= bottom; a-- {
		for b := top; b >= bottom; b-- {
			tmp = a * b
			if isPalindrome(tmp) && tmp > max {
				max = tmp
			}
		}
	}
	fmt.Println(time.Now())
	fmt.Println(max)
}

func isPalindrome(num int) bool {
	b := strconv.Itoa(num)
	return b == reverse2(b)
}

func reverse(s string) string {
	if len(s) <= 1 {
		return s
	}

	arr := strings.Split(s, "")
	length := len(arr)
	halfLen := length / 2
	for i := 0; i < halfLen; i++ {
		arr[i], arr[length-i-1] = arr[length-i-1], arr[i]
	}
	return strings.Join(arr, "")
}

func reverse2(s string) string {
	if len(s) <= 1 {
		return s
	}

	r := []rune(s)
	length := len(r)
	halfLen := length / 2
	for i := 0; i < halfLen; i++ {
		r[i], r[length-i-1] = r[length-i-1], r[i]
	}
	return string(r)
}
