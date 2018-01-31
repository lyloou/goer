package main

import (
	"reflect"
	"fmt"
)

func Method(in interface{}) (ok bool) {
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Slice {
		ok = true
	} else {
		//panic
	}

	num := v.Len()
	for i := 0; i < num; i++ {
		i := v.Index(i).Interface()

		fmt.Printf("%T\n", i)
	}
	return ok
}

func main() {
	s := []int{1, 3, 5, 7, 9}
	b := []float64{1.2, 3.4, 5.6, 7.8}
	Method(s)
	Method(b)
}
