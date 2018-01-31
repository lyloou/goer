package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%x", buf))
}
