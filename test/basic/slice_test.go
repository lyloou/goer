package basic

import (
	"fmt"
	"strings"
	"testing"
)

type Retriever struct {
	Contents string
}

func (r Retriever) String() string {
	return r.Contents
}

func TestSlice(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("arr=%v, len(arr)=%v, cap(arr)=%v\n", arr, len(arr), cap(arr))

	s1 := arr[4:6]
	fmt.Printf("s1=%v, len(s1)=%v, cap(s1)=%v\n", s1, len(s1), cap(s1))
	s2 := s1[:3]
	fmt.Printf("s2=%v, len(s2)=%v, cap(s2)=%v\n", s2, len(s2), cap(s2))

	arr[4] = 1001
	s2[2] = 100
	fmt.Println(arr)
	fmt.Println(s2)
	s2 = append(s2, 101)
	s2 = append(s2, 102)
	s2 = append(s2, 103)
	s2 = append(s2, 104)
	fmt.Println(arr)
	fmt.Println(s2)
	arr[7] = 1000
	arr[4] = 1008
	fmt.Println(arr)
	fmt.Println(s2)

}

func TestSlice2(t *testing.T) {
	b := make([]byte, 10)
	s := "nihao"
	fmt.Println(b)
	fmt.Println(s)
	fmt.Println(strings.NewReader(s).Read(b))
	fmt.Println(b)
}

func TestSlice3(t *testing.T) {
	b := []int{1}
	a := b[0]
	b = b[1:]
	fmt.Println(a, b)

}