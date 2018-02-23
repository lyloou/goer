package main

import "fmt"

const theNum = 1000

func main() {
outer:
	for i := 1; i < theNum; i++ {
		for j := i + 1; j < theNum; j++ {
			for k := j + 1; k < theNum; k++ {
				if i*i+j*j == k*k {
					fmt.Printf("i:%d, j:%d, k:%d\n", i, j, k)
					if i+j+k == theNum {
						fmt.Printf("got it:%d, %d, %d\n", i, j, k)
						fmt.Printf("result:%d\n", i*j*k)
						break outer
					}
				}
			}
		}
	}
}
