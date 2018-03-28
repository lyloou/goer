package main

import (
	"fmt"
	"strings"

	"github.com/lyloou/goer/cmd/project-euler/common"
)

// 分析： 对于被乘数，乘数，积，按长度来算，它们三个的长度和必须为9

func main() {
	sum := getPandigitalProductsSum()
	fmt.Println(sum)
}

func getPandigitalProductsSum() int {
	// 使用map的目的是去重
	productMap := make(map[int]int, 0)
	// len(i) = 1, len(j) == 4, len(product) = 4
	for i := 1; i < 10; i++ {
		for j := 1000; j < 10000; j++ {
			if !canUse(i, j) {
				continue
			}
			if product, ok := getPandigitalProduct(i, j); ok {
				productMap[product] = product
				fmt.Println(i, j, product)
			}
		}
	}
	// len(i) = 2, len(j) == 3, len(product) = 4
	for i := 10; i < 100; i++ {
		for j := 100; j < 1000; j++ {
			if !canUse(i, j) {
				continue
			}
			if product, ok := getPandigitalProduct(i, j); ok {
				productMap[product] = product
				fmt.Println(i, j, product)
			}
		}
	}
	sum := 0
	for _, v := range productMap {
		sum += v
	}
	return sum
}

func canUse(i int, j int) bool {
	a0 := fmt.Sprintf("%d", 0)
	ai := fmt.Sprintf("%d", i)
	aj := fmt.Sprintf("%d", j)

	return !strings.ContainsAny(ai, a0) &&
		!strings.ContainsAny(aj, a0) &&
		!strings.ContainsAny(ai, aj) &&
		!strings.ContainsAny(aj, ai) &&
		common.JustShowOnce(ai) &&
		common.JustShowOnce(aj)
}

func getPandigitalProduct(num1, num2 int) (product int, ok bool) {
	product = num1 * num2
	return product, isPandigital(num1, num2, product)
}

func isPandigital(num1 int, num2 int, num3 int) bool {
	a0 := fmt.Sprintf("%d", 0)

	s := fmt.Sprintf("%d%d%d", num1, num2, num3)
	if len(s) != 9 {
		return false
	}
	if strings.ContainsAny(s, a0) {
		return false
	}
	if !common.JustShowOnce(s) {
		return false
	}

	return true
}
