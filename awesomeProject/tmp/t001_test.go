package main

import (
	"testing"
)

func TestOk(t *testing.T) {
	t.Log("kkkkkkk")

}

func test() string {
	s := make([] string, 1000)
	for i := 0; i < 1000; i++ {
		s[i] = "a"
	}
	return Join(s, "")
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}

func Join(a []string, sep string) string {
	n := len(sep) * (len(a) - 1)

	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}

	return string(b)
}
