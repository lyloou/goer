package main

import "fmt"

func main() {

}


func sample1() {
	fmt.Println("Ok, now!")
	var tmp int64 = 32
	var c interface{} = tmp
	switch c.(type) {
	case int:
		fmt.Println("int:", c.(int))
	case int64:
		fmt.Println("int64:", c.(int64))
	case string:
		fmt.Println("string:", c.(string))
	case float64:
		fmt.Println("float:", c.(float64))
	}
	fmt.Println("c", c)
	fmt.Printf("===>%x", c)
	fmt.Printf("==>%x", &c)
	fmt.Println("==>", 2<<3)
}
