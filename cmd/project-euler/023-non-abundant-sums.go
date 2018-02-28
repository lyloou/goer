package main

import (
	"fmt"

	"github.com/lyloou/goer/cmd/project-euler/common"
	"github.com/lyloou/goer/pkg/timing"
)

// https://www.mathblog.dk/project-euler-23-find-positive-integers-not-sum-of-abundant-numbers/

func main() {
	start := timing.Start()
	var limit int64 = 28213
	abundant := make([]int64, 0)
	var i, j int64
	for i = 1; i <= limit; i++ {
		if common.SumDivisors(common.GetDivisors(i)) > i {
			abundant = append(abundant, i)
		}
	}
	fmt.Println(timing.Stop(start))
	fmt.Println(abundant)

	canBeWrittenAsAbundant := make([]bool, limit+1)
	length := int64(len(abundant))
	for i = 0; i < length; i++ {
		for j = i; j < length; j++ {
			canBeWritten := abundant[i] + abundant[j]
			if canBeWritten <= limit {
				canBeWrittenAsAbundant[canBeWritten] = true
			}
		}
	}
	fmt.Println(canBeWrittenAsAbundant)

	var result int64
	for i = 1; i <= limit; i++ {
		if !canBeWrittenAsAbundant[i] {
			result += i
		}
	}
	fmt.Println(result)

}
