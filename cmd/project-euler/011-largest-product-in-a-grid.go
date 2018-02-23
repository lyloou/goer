package main

import (
	"github.com/lyloou/goer/cmd/project-euler/handler"
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

const length = 20
const num = 4

func main() {
	// save to datas
	datas := make([][]int, length)
	err := handler.HandleScannerWithFilename("011-largest-product-in-a-grid.data",
		func(scanner *bufio.Scanner) error {
			for i := 0; i < len(datas); i++ {
				datas[i] = make([]int, length)
			}

			x := 0
			for scanner.Scan() {
				s := strings.Split(scanner.Text(), " ")
				for k, v := range s {
					tmp, err := strconv.Atoi(v)
					if err != nil {
						return err
					}
					datas[x][k] = tmp
				}
				x++
			}

			return nil
		})

	if err != nil {
		panic(err)
	}

	max := 0
	// right
	fmt.Println("left -> right")
	for i := 0; i < length; i++ {
		for j := 0; j <= length-num; j++ {
			result := datas[i][j] * datas[i][j+1] * datas[i][j+2] * datas[i][j+3]
			if result > max {
				max = result
				fmt.Printf("%8d = %02d(%02d, %02d) * %02d(%02d, %02d) * %02d(%02d, %02d) * %02d(%02d, %02d)\n",
					max,
					datas[i][j], i, j,
					datas[i][j+1], i, j+1,
					datas[i][j+2], i, j+2,
					datas[i][j+3], i, j+3)
			}
		}
	}

	// down
	fmt.Println("up -> down")
	for i := 0; i <= length-num; i++ {
		for j := 0; j < length; j++ {
			result := datas[i][j] * datas[i+1][j] * datas[i+2][j] * datas[i+3][j]
			if result > max {
				max = result
				fmt.Printf("%8d = %02d(%02d, %02d) * %02d(%02d, %02d) * %02d(%02d, %02d) * %02d(%02d, %02d)\n",
					max,
					datas[i][j], i, j,
					datas[i+1][j], i+1, j,
					datas[i+2][j], i+2, j,
					datas[i+3][j], i+3, j)
			}
		}
	}

	// diagonally
	fmt.Println("left-up -> right-down")
	for i := 0; i <= length-num; i++ {
		for j := 0; j <= length-num; j++ {
			result := datas[i][j] * datas[i+1][j+1] * datas[i+2][j+2] * datas[i+3][j+3]
			if result > max {
				max = result
				fmt.Printf("%8d = %02d(%02d, %02d) * %02d(%02d, %02d) * %02d(%02d, %02d) * %02d(%02d, %02d)\n",
					max,
					datas[i][j], i, j,
					datas[i+1][j+1], i+1, j+1,
					datas[i+2][j+2], i+2, j+2,
					datas[i+3][j+3], i+3, j+3)
			}
		}
	}

	// diagonally
	fmt.Println("left-down -> right-up")
	for i := length - num; i >= num; i-- {
		for j := 0; j <= length-num; j++ {
			result := datas[i][j] * datas[i-1][j+1] * datas[i-2][j+2] * datas[i-3][j+3]
			if result > max {
				max = result
				fmt.Printf("%8d = %02d(%02d, %02d) * %02d(%02d, %02d) * %02d(%02d, %02d) * %02d(%02d, %02d)\n",
					max,
					datas[i][j], i, j,
					datas[i-1][j+1], i-1, j+1,
					datas[i-2][j+2], i-2, j+2,
					datas[i-3][j+3], i-3, j+3)
			}
		}
	}
}
