package main

import (
	"net/http"
	"fmt"
	"time"
)

func main(){
	fmt.Println(time.Now())
	resp, err := http.Head("http://google.com")
	fmt.Println(resp, err)
	fmt.Println(time.Now())
}
