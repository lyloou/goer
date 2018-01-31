
## [Go – taking slices of any type as input parameters](https://ahmet.im/blog/golang-take-slices-of-any-type-as-input-parameter/)
```go
package slice

func takeSliceArg(arg interface{}) (out []interface{}, ok bool) {
    slice, success := takeArg(arg, reflect.Slice)
    if !success {
        ok = false
          return
    }
    c := slice.Len()
    out = make([]interface{}, c)
    for i := 0; i < c; i++ {
        out[i] = slice.Index(i).Interface()
    }
    return out, true
}

func takeArg(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
    val = reflect.ValueOf(arg)
    if val.Kind() == kind {
        ok = true
    }
    return
}
```

## [谈一谈Go的interface和reflect | Legendtkl](http://legendtkl.com/2015/11/28/go-interface-reflect/)
## [collection-functions](https://gobyexample.com/collection-functions)

## [Removing item(s) from a slice, while iterating in Go](https://vbauerster.github.io/2017/04/removing-items-from-a-slice-while-iterating-in-go/)
```go
package main

import "fmt"

func main() {
	arr := []int64{1, 2, 3, 4, 6, 8, 9, 10, 11, 23}
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i]%2 == 0 {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	fmt.Println(arr)
}
```