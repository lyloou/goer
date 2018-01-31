package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
	name string
	age  int
}

func main() {
	fmt.Printf("%#v\n", User{
		name: "nihao",
		age:  32,
	})

	switch "dm" {
	case "dm":
		fmt.Println("dddd")
	case "dm2":
		fmt.Println("dddd2222")
	}

	a := 5
	switch  {
	case a == 3:
		fmt.Println(">=4")
		fallthrough
	case a == 5:
		//fmt.Println(">=5")
		fallthrough
	default:
		fmt.Println("default")
	}
	aa := `{"products":null,"spec":null}`
	json.Unmarshal([]byte(aa), &aa)


}

