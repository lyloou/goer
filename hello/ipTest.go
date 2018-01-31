package main

import (
	"fmt"
	"time"
	"io"
	"crypto/md5"
	"regexp"
)

type Ok int

func main() {
	var v interface{}
	v = Ok(23)
	o, i := v.(int)

	fmt.Println(o, i)

	tim := time.Now().String()
	fmt.Println(tim)
	io.WriteString(md5.New(), tim)

	fmt.Printf("hello, world\n") // hello
	fmt.Printf("hello\n\n")      // world
	fmt.Println(IsIp("192.168.0.10"))
}

func IsIp(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}
	return true
}
