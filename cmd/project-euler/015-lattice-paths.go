package main

import (
	"fmt"
)

func main() {
	// 和method1方法一样，但是数值太大了，溢出了，所以要用化简后的方法来计算。
	const grid int64 = 20
	var j, product int64 = 1, 1
	for i := grid + 1; i <= 2*grid; i++ {
		product = product * i / j
		j++
	}
	fmt.Println(product)

	countRouters(20, 20)
}

// http://blog.csdn.net/liangzhaoyang1/article/details/52939462
// https://zh.wikipedia.org/wiki/%E7%BB%84%E5%90%88%E6%95%B0%E5%AD%A6#%E6%80%BB%E7%BB%93
// https://baike.baidu.com/item/%E6%8E%92%E5%88%97%E7%BB%84%E5%90%88/706498
// https://gist.github.com/aleksa/a90f38bfe4d6af163358

// 组合的方式
func method1(grid int64) int64 {
	return getFactorial(grid*2) / (getFactorial(grid) * getFactorial(grid))
}

func getFactorial(num int64) int64 {
	var result, i int64 = 1, 1
	for ; i <= num; i++ {
		result *= i
	}
	return result
}

type point struct {
	i, j int
}

func (p point) add(pp point) point {
	return point{p.i + pp.i, p.j + pp.j}
}

var dir = [2]point{
	{0, 1},
	{1, 0},
}

var sum = 0

// 递归的方式，求20x20行不通（<=16 就有些够呛）；
func method2() {
	grid := 4
	start := point{0, 0}
	end := point{grid, grid}
	walkToNext(start, end)
	//walkToNext2(start, end)
	fmt.Println(sum)
}

func walkToNext(p, e point) {

	for _, d := range dir {
		pp := p
		pp = pp.add(d)
		if pp == e {
			sum += 1
		}

		if less(pp, e) {
			walkToNext(pp, e)
		}
	}
}

func less(p point, e point) bool {
	return !(p.i > e.i || p.j > e.j)
}

func walkToNext2(p, e point) {

	for _, d := range dir {
		pp := p
		pp = pp.add(d)

		if pp.i <= e.i && pp.j <= e.j {
			fmt.Println(pp)
			if pp == e {
				sum += 1
				fmt.Println()
			} else {
				walkToNext(pp, e)
			}
		}
	}
}

// https://projecteuler.net/overview=015
// 迭代的方式
func countRouters(m, n int) {
	grid := make([][]int, m+1)
	for k, _ := range grid {
		grid[k] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		grid[i][0] = 1
	}

	for i := 0; i <= m; i++ {
		grid[0][i] = 1
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			fmt.Printf("%12d ", grid[i][j])
		}
		fmt.Printf("\n")
	}

}
