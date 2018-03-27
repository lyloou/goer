package main

import (
	"bytes"
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")
var c = flag.Int("lou", 1, "input lou")

func main() {
	flag.Parse()
	fmt.Println(*c)

	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println("\nend")

	var value float64
	var unit string
	fmt.Sscanf("12.0332cC", "%.2f%s", &value, &unit)
	fmt.Println(value, unit)

	var buf *bytes.Buffer
	buf = new(bytes.Buffer) // buf 变量赋予了一个*bytes.Buffer的空指针
	fmt.Println(buf == nil) // false。 buf的动态类型和动态值必须同时为nil，该结果才为true;
	// https://yar999.gitbooks.io/gopl-zh/content/ch7/ch7-05.html
	fmt.Println(buf.WriteRune(12423))
	fmt.Println(buf)

	double := func(x int) (result int) {
		defer func() {
			fmt.Printf("double(%d) = %d\n", x, result)
			result = 18
		}()
		return x + x
	}
	fmt.Println(double(4))
}
