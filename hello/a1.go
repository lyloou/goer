package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"reflect"
	"path/filepath"
	"flag"
	"encoding/base64"
)

func main() {
	const intSize = 32 << (^uint(0) >> 63)
	fmt.Println(intSize)

	encoding := base64.StdEncoding.EncodeToString([]byte("This train bounds for JinZhou."))
	fmt.Println(encoding)
	decoding, _ := base64.StdEncoding.DecodeString(encoding)
	fmt.Println(string(decoding))
}

func sampl7() {
	var d = `
	ok

	ok2

	ok3
	`
	split := strings.Split(d, "\n")
	fmt.Println(len(split))
	fmt.Println(split)
	fmt.Println(os.Getenv("GOROOT"))
	fmt.Println(os.Getenv("PATH"))
	fmt.Println(*flag.String("addr", ":1718", "http service address"))
}
func sampl6() {
	fmt.Println(filepath.ToSlash(`{"addr":"192.168.0.23:6379","password":"123456","db":0,"pool_size":100}`))
	var user User
	user.Id = "1"
	user.Name = "user"
	fmt.Println(user)
	u := new(User)
	u.Id = "2"
	u.Name = "user2"
	fmt.Println(u)
}

type User struct {
	Id   string
	Name string
}

func sampl5() {
	fmt.Println("abc" == "abC")
	v := "k"
	var vv interface{}
	vv, _ = strconv.ParseInt(v, 16, 64)
	i, ok := vv.(int64)
	av, _ := strconv.Atoi(v)
	fmt.Println(i, ok, reflect.TypeOf(vv), reflect.TypeOf(av))
	fmt.Println(int(float64(32)))
}
func sampl4() {
	println(strings.TrimRight("ok/", "/"))
}

func sampl3() {
	f, err := os.Open("ok.txt")
	if err != nil {
		fmt.Println("", err.Error())
		return
	}
	defer f.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := f.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
func smaple1() {
	a, b := 1, 2
	a, b = b, a
	// 交换，先计算右边的；再赋值给左边；
	fmt.Println(a, b)
}
