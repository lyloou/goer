package main

import (
	"time"
	"fmt"
	"github.com/lyloou/goer/cmd/project-euler/common"
)

func main() {
	m019of2()
}

func m019of2() {
	fmt.Println(common.IsLeap(2001))

	fmt.Println(theDaysFromYourBirth(time.Date(1992, time.Month(11), 16, 0, 0, 0, 0, time.Local)))

	// 计算每月 1 号距离 1900.1.1 日的天数，取余和 Sunday 比较
	sum := 0
	for y := 1901; y <= 2000; y++ {
		for m := 1; m <= 12; m++ {
			date := time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.Local)
			if isSunday(date) {
				sum ++
			}
		}
	}
	fmt.Println(sum)
}

func isSunday(date time.Time) bool {
	var date1 = time.Date(1900, time.Month(1), 7, 0, 0, 0, 0, time.Local)
	return differDays(date, date1)%7 == 0
}

func differDays(time1, time2 time.Time) int64 {
	differ := time1.UnixNano() - time2.UnixNano()
	if differ < 0 {
		differ = -differ
	}
	day := differ / (time.Hour.Nanoseconds() * 24)
	return day
}

func theDaysFromYourBirth(birth time.Time) int64 {
	return differDays(birth, time.Now())
}

func m019of1() {
	sum := 0
	for y := 1901; y <= 2000; y++ {
		for m := 1; m <= 12; m++ {
			date := time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.Local)
			if date.Weekday() == time.Sunday {
				sum ++
			}
		}
	}
	fmt.Println(sum)
}
