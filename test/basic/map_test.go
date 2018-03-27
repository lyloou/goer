package basic

import (
	"fmt"
	"net/url"
	"testing"
)

// https://yar999.gitbooks.io/gopl-zh/content/ch6/ch6-02.html
func TestMap(t *testing.T) {
	m := url.Values{"lang": {"en"}} // direct construction
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang")) // "en"
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // "1"      (first value)
	fmt.Println(m["item"])     // "[1 2]"  (direct map access)
	fmt.Println(m.Get("item")) // "1"      (first value)

	m.Set("item", "3")
	fmt.Println("set:", m.Get("item")) // "3"
	fmt.Println(m["item"])     // "3"
	m.Add("item", "4")
	fmt.Println(m["item"])     // "3"

	m.Del("item")
	fmt.Printf("del:%q\n", m.Get("item")) // del:""

	m = nil
	fmt.Println(m.Get("item")) // ""
	//m.Add("item", "3")         // panic: assignment to entry in nil map
}
