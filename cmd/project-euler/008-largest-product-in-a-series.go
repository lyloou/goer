package main

import (
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"github.com/lyloou/goer/cmd/project-euler/handler"
)

const count = 1000
const step = 13

func main() {
	var i int64 = 0
	numbers := make([]int64, count)

	err := handler.HandleScannerWithFilename("008-largest-product-in-a-series.data",
		func(scanner *bufio.Scanner) error {
			// 转换为长数组
			for scanner.Scan() {
				str := strings.Split(scanner.Text(), "")
				for _, v := range str {
					n, err := strconv.Atoi(v)
					if err != nil {
						return err
					}
					numbers[i] = int64(n)
					i++
				}
			}
			return nil
		})

	if err != nil {
		panic(err)
		return
	}

	i = 0
	var max int64 = 0
	var theArray []int64
	for ; i < count-step; i++ {
		result, number := getProduct(numbers, i)
		fmt.Printf("number:%v, result:%v \n", number, result)

		// 与之前的最大乘积作比较
		if result > max {
			max = result
			theArray = number
		}
	}

	fmt.Printf("\ntheArray:%v, max:%v\n", theArray, max)
}

// 获取这个step位数字的乘积
func getProduct(numbers []int64, index int64) (result int64, number []int64) {
	// 取一个step位数字
	number = make([]int64, step)
	result = 1
	var i int64
	for i = 0; i < step; i++ {
		number[i] = numbers[index+i]
		result = result * number[i]
	}
	return
}
