package method

import (
	"fmt"
	"reflect"
	"strings"
)

// https://yar999.gitbooks.io/gopl-zh/content/ch12/ch12-08.html
func Print(x interface{}){
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s", t)

	for i:= 0 ; i<v.NumMethod(); i++{
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name, strings.TrimPrefix(methType.String(), "func"))
	}
}