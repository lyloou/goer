package main

import (
	"fmt"
	"strings"
)

var onThousand = "one thousand"
var lessThan20Array = []string{"", "one", "two", "three", "four",
	"five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve",
	"thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var lessThan20 = make(map[int]string, len(lessThan20Array))

var tenTimesArray = []string{"twenty", "thirty", "forty", "fifty",
	"sixty", "seventy", "eighty", "ninety",}
var tenTimes = make(map[int]string, len(tenTimesArray))

func init() {
	for i := 0; i < 20; i++ {
		lessThan20[i] = lessThan20Array[i]
	}

	for i := 20; i < 100; i += 10 {
		tenTimes[i] = tenTimesArray[(i-20)/10]
	}
}

func getLessThan20(num int) string {
	return lessThan20[num]
}

func getLess100(num int) string {
	if num < 20 {
		return getLessThan20(num)
	}

	times := num / 10 * 10
	return fmt.Sprintf("%s-%s", tenTimes[times], lessThan20[num%times])
}

func getLess999(num int) string {
	if num < 100 {
		return getLess100(num)
	}
	h := num / 100
	result := ""
	result += lessThan20[h] + " hundred"
	if t := num % 100; t != 0 {
		result += " and " + getLess100(t)
	}
	return result
}

func main() {
	sum := 0
	for i := 1; i < 1000; i++ {
		sum += getValidLettersNumber(getLess999(i))
	}
	sum += getValidLettersNumber(onThousand)
	fmt.Println(sum)
}

func getValidLettersNumber(num string) int {
	count := 0
	str := strings.Split(num, "")
	for _, v := range str {
		if " " != v && "-" != v {
			count ++
		}
	}
	return count
}
