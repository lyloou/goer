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

func TestSlice4Append(t *testing.T) {
	data := []string{"A", "B", "C"}
	t.Log(data)
	data = append(data, "D")
	t.Log(data)
}

// https://codingair.wordpress.com/2014/07/18/go-appendprepend-item-into-slice/
// https://code.google.com/p/go-wiki/wiki/SliceTricks
func TestSlice5Prepend(t *testing.T) {
	data := []string{"A", "B", "C"}
	t.Log(data)
	data = append([]string{"D"}, data...)
	t.Log(data)
}

// https://yar999.gitbooks.io/gopl-zh/content/ch4/ch4-02.html
func TestSliceModifyArrayWillChangeSlice(t *testing.T) {
	months := [13]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Printf("%T\n", months)
	Q2 := months[4:7]
	summer := months[6:9]
	months[6] = 87
	fmt.Println(Q2, summer)
	Q2[len(Q2)-1] = 88
	fmt.Println(Q2, summer)
	summer[0] = 89
	fmt.Println(Q2, summer)
}
