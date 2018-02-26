package main

import (
	"fmt"
	"strings"
)

// https://www.mathsisfun.com/numbers/counting-names-1000.html
func main() {

	sum := 0
	for i := 1; i < 1000; i++ {
		sum += getLettersNumber(less999(i))
	}
	sum += getLettersNumber(equal1000())
	fmt.Println(sum)
}

func getLettersNumber(num string) int {
	num = strings.Replace(num, "-", "", -1)
	num = strings.Replace(num, " ", "", -1)
	//num = strings.Join(strings.Split(num, "-"), "")
	//num = strings.Join(strings.Split(num, " "), "")
	return len(num)
}

func less20(num int) string {
	switch num {
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	case 10:
		return "ten"
	case 11:
		return "eleven"
	case 12:
		return "twelve"
	case 13:
		return "thirteen"
	case 14:
		return "fourteen"
	case 15:
		return "fifteen"
	case 16:
		return "sixteen"
	case 17:
		return "seventeen"
	case 18:
		return "eighteen"
	case 19:
		return "nineteen"
	default:
		return ""
	}

}

func less100(num int) string {
	switch {
	case num <= 19:
		return less20(num)
	case num <= 29:
		return "twenty-" + less20(num%20)
	case num <= 39:
		return "thirty-" + less20(num%30)
	case num <= 49:
		return "forty-" + less20(num%40)
	case num <= 59:
		return "fifty-" + less20(num%50)
	case num <= 69:
		return "sixty-" + less20(num%60)
	case num <= 79:
		return "seventy-" + less20(num%70)
	case num <= 89:
		return "eighty-" + less20(num%80)
	case num <= 99:
		return "ninety-" + less20(num%90)
	default:
		return ""
	}

}

func less999(num int) string {
	if num < 100 {
		return less100(num)
	}

	h := num / 100
	result := ""
	result += less20(h) + " hundred"
	if t := num % 100; t != 0 {
		result += " and " + less100(t)
	}
	return result
}

func equal1000() string {
	return "one thousand"
}
