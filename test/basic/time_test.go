package basic

import (
	"testing"
	"time"
	"fmt"
)

// https://stackoverflow.com/questions/5885486/how-do-i-get-the-current-timestamp-in-go
func TestTimestamp(t *testing.T) {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Format("20060102150405"))
	fmt.Println(time.Now().Format(time.RFC850))
	fmt.Println(time.Now().Format(time.RFC822))
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}

func TestAfter(t *testing.T) {
	fmt.Println(time.Now())
	<-time.After(1e3 * time.Millisecond) // 1000ms
	fmt.Println(time.Now())
	<-time.After(1e9) // 1s
	fmt.Println(time.Now())
}
