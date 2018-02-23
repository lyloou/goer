package main

import (
	"fmt"
	"github.com/lyloou/goer/cmd/project-euler/common"
)

// http://zhex.iteye.com/blog/489729
func main() {
	//var data, index int64 = 60, 2 // => [2 2 3 5]
	//var data, index int64 = 13195, 2 // => [5 7 13 29]
	var data, index int64 = 600851475143, 2 // => [71 839 1471 6857]
	var arr []int64
	for index <= data {
		if common.IsPrime(index) {
			if data%index == 0 {
				data = data / index
				arr = append(arr, index)
			} else {
				index ++
			}
		} else {
			index ++
		}
	}
	fmt.Println(arr)
}
