package main

import (
	"fmt"
)

var lessThan20Array = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var lessThan20 = make(map[int]string, len(lessThan20Array))

var tenTimesArray = []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",}
var tenTimes = make(map[int]string, len(tenTimesArray))

func init() {
	for i := 0; i < 20; i++ {
		lessThan20[i] = lessThan20Array[i]
	}

	for i := 20; i < 100; i += 10 {
		tenTimes[i] = tenTimesArray[(i-20)/10]
	}
}

// https://www.mathsisfun.com/numbers/counting-names-1000.html
func main() {
	fmt.Println(lessThan20)
	fmt.Println(tenTimes)

	for i := 1; i < 1000; i++ {
		if i < 100 {
			if i <= 19 {
				fmt.Println(lessThan20[i])
				continue
			}
			times := i / 10 * 10
			fmt.Printf("%s-%s\n", tenTimes[times], lessThan20[i%times])
			continue
		}

		h := i / 100
		result := ""
		result += lessThan20[h] + " hundred"
		if t := i % 100; t != 0 {
			if t <= 19 {
				result += " and " + lessThan20[t]
				continue
			}
			times := t / 10 * 10
			result += " and " + tenTimes[times] + "-" + lessThan20[t%times]
		}
		fmt.Println(result)
	}

}
