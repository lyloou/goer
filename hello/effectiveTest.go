package main

import (
	"fmt"
	"os"
)

func main() {
	var x uint64 = 1<<64-1
	fmt.Println(x)
	fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))
}


func sampleB() {
	fmt.Println(fmt.Sprintf("kek ksdfj %d woekk dkfj", 2))
	fmt.Fprintf(os.Stderr, "%s dkfj %d", "wowowo", 32)
	var c = struct {
		name string
		age  int
	}{
		name: "kd",
		age:  23,
	}
	fmt.Println(c)
}


func sampleA() {
	attended := map[string]int{
		"Ann": 1,
		"Joe": 2,
		"Bob": 3,
	}
	i, k := attended["Bob"]
	fmt.Println(i, k)
}


func sample9() {
	// f := os.NewFile(777, "ooo")
	b := make([]string, 0, 10)
	strings := append(b, "kjasldkfaslkdfalskdfjkasdfjaskdl")
	fmt.Println(strings)
}

func sample8() {
	var s = "ok ksdjfks skdlfjiwi sdkfjkwl"
	for v, a := range s {
		fmt.Println(v, string(a))
	}
}

func sample7() {
	file := os.NewFile(777, "ok.txt")
	defer file.Close()
	fmt.Println(new(os.File) == &os.File{})
	fmt.Println(file.Fd())
	fmt.Println(file.Stat())
	a := [...]string{"no error", "Eio", "invalid argument"}
	s := []string{"no error", "Eio", "invalid argument"}
	fmt.Println(a)
	fmt.Println(s)
	changeStr(&s)
	fmt.Println(s)
	c := new(int)
	*c = 32
	fmt.Println(*c)
	var v = make([]byte, 10, 20)
	v[0] = 0
	v[1] = 1
	v[2] = 2
	v[9] = 32
	v = v[:cap(v)]
	// We can grow s to its capacity by slicing it again  https://blog.golang.org/go-slices-usage-and-internals
	v[10] = 33
	fmt.Println(v)
	var n = new([]byte)
	n = &v
	fmt.Println(*n)
}
func changeStr(str *[]string) {
	fmt.Println(len(*str))
	(*str)[0] = "ksdaf"
	(*str)[1] = "ksdaf"
	(*str)[2] = "ksdaf"
}

func sample6() {
	var e error
	var na Na
	var ia Ia
	fmt.Println(na, e, ia)
	n := new(Na)
	i := new(Ia)
	fmt.Println(n, i)
	c := make([]byte, 2, 5)
	fmt.Println(c)
}

type Na struct {
	name string
}

type Ia interface {
}

func sample5() {
	bbs := []byte{8, 'a', 2, 'o', 'b', 'd', 'c', 3, 2, '1', 1}
	x := 0
	for i := 0; i < len(bbs); i++ {
		x, i = nextInt(bbs, i)
		fmt.Println(x, i)
	}
	bbs = []byte("hillo")
	fmt.Println(string(bbs))
}

func sample4() {
	var of os.File
	f, err := of.Write([]byte("ok leee"))
	fmt.Println(f, err)
	fmt.Println(multiple(2))
}

func multiple(i int) (a int, b int, c error) {
	a = i
	return a, b, c
}
func samplee3() {
	var t interface{}
	c := true
	t = &c
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}
}
func samplee2() (int, error) {
	return fmt.Println(Compare([]byte{1, 2, 3, 4}, []byte{1, 2, 3}))
}
func samplee1() {
	for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// Compare returns an integer comparing the two byte slices,
// lexicographically.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b
func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}

func nextInt(b []byte, i int) (int, int) {
	for ; i < len(b) && !isDigit(b[i]); i++ {
	}
	x := 0
	for ; i < len(b) && isDigit(b[i]); i++ {
		x = x*10 + int(b[i])
	}
	return x, i
}
func isDigit(b byte) bool {
	return b >= 0 && b <= 9
}
