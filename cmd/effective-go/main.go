package main

import (
	"fmt"
	"encoding/json"
	"strings"
	"github.com/pkg/errors"
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
	switch {
	case a == 3:
		fmt.Println(">=4")
		fallthrough
	case a == 5:
		//fmt.Println(">=5")
		fallthrough
	default:
		fmt.Println("default")
	}

	type goods struct {
		Products string
		Spec     string
	}

	aa := `{"products":"knife","spec":"small"}`
	g := goods{}
	json.Unmarshal([]byte(aa), &g)
	fmt.Printf("%+v\n", g)

	c := make([]byte, 2)
	strings.NewReader("ok").Read(c)
	fmt.Printf("%s", c)

	errors.New("ldk")

}
