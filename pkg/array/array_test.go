package array

import (
	"fmt"
	"strings"
	"testing"
)

func TestInt64(t *testing.T) {

	srcInt64Arr := []int64{1, 2, 3, 4, 6, 8, 9, 10, 11, 23}
	filterInt64Arr := FilterFromInt64Array(srcInt64Arr, isEven)
	removeInt64Arr := RemoveFromInt64Array(srcInt64Arr, func(item interface{}) bool {
		v, ok := item.(int64)
		if ok {
			return v > 9
		}
		return false
	})
	fmt.Println("srcInt64Arr", srcInt64Arr)
	fmt.Println("filterInt64Arr", filterInt64Arr)
	fmt.Println("removeInt64Arr", removeInt64Arr)
}

func isEven(number interface{}) bool {
	v, ok := number.(int64)
	if ok {
		return v%2 == 0
	}
	return false
}

func TestString(t *testing.T) {
	srcStringArr := []string{"ok1", "baibai", "ok2", "nihao", "whoareyou", "ok3"}
	filterStringArr := FilterFromStringArray(srcStringArr, startWith) // will not change srcStringArr
	removeStringArr := RemoveFromStringArray(srcStringArr, func(item interface{}) bool {
		v, ok := item.(string)
		if ok {
			return strings.EqualFold(v, "whoareyou")
		}
		return false
	})
	fmt.Println("srcStringArr", srcStringArr)
	fmt.Println("filterStringArr", filterStringArr)
	fmt.Println("removeStringArr", removeStringArr)

	b := make([]string, 0) // will not change srcStringArr
	for _, x := range srcStringArr {
		if startWith(x) {
			b = append(b, x)
		}
	}
	fmt.Println(b)
	fmt.Println(srcStringArr)

	// [Can I convert a []T to an []interface{}?](https://golang.org/doc/faq#convert_slice_of_interface)
	c := srcStringArr[:0] // !!! will change srcStringArr
	for _, x := range srcStringArr {
		if startWith(x) {
			c = append(c, x)
		}
	}
	fmt.Println(c)
	fmt.Println(srcStringArr)
}

func startWith(str interface{}) bool {
	v, ok := str.(string)
	if ok {
		return strings.HasPrefix(v, "ok")
	}
	return false
}
