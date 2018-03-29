package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	numerator, denominator := 1, 1
	for i := 10; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			if isDigitCancellingFraction(i, j) {
				fmt.Println("got:", i, j)
				numerator *= i
				denominator *= j
				break
			}
		}
	}
	fmt.Println(float64(denominator) / float64(numerator))
}

func isDigitCancellingFraction(num1, num2 int) bool {
	strNum1 := fmt.Sprintf("%d", num1)
	strNum2 := fmt.Sprintf("%d", num2)

	if strings.Contains(strNum1, "0") {
		return false
	}

	for _, v := range strNum1 {
		vv := fmt.Sprintf("%c", v)
		if strings.Contains(strNum2, vv) {
			cNum1 := strings.Replace(strNum1, vv, "", 1)
			cNum2 := strings.Replace(strNum2, vv, "", 1)
			n1, _ := strconv.Atoi(cNum1)
			n2, _ := strconv.Atoi(cNum2)
			if float32(n1)/float32(n2) == float32(num1)/float32(num2) {
				return true
			}
		}
	}
	return false
}
