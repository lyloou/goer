package main

import (
	"os"
	"fmt"
	"log"
)

type name string

func (n *name)prname(str string)  {
	fmt.Println(str)
	fmt.Println(*n)
}

func main() {
	var d name
	d = "okdk"
	d.prname("wowowo")
	fmt.Println(len(d))

	file, err := os.Open("errorTest.go")
	if err != nil {
		log.Fatal(err)
		return
	}
	f := make([]byte, 1024)
	fmt.Println(file.Read(f))
	fmt.Println(string(f))
}
