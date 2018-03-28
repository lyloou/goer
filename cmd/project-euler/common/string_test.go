package common

import (
	"testing"
)

func TestJustShowOnce(t *testing.T) {
	tests := []struct {
		str  string
		want bool
	}{
		{"1", true},
		{"a", true},
		{"abce", true},
		{"1234", true},

		{"1214", false},
		{"abab", false},
		{"你好啊你", false},
	}

	for _, v := range tests {
		if got := JustShowOnce(v.str); got != v.want {
			t.Errorf("JustShowOnce(%s) got:%v, expected:%v", v.str, got, v.want)
		}
	}
}

func BenchmarkJustShowOnce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JustShowOnce("你好啊你")
	}
}

func TestContainAll(t *testing.T) {
	tests := []struct {
		str1 string
		str2 string
		want bool
	}{
		{"1214", "1124", true},
		{"你好啊你", "你好你啊", true},

		{"abab", "aabc", false},
		{"123", "12", false},
	}

	for _, v := range tests {
		if got := IsPermutation(v.str1, v.str2); got != v.want {
			t.Errorf("IsPermutation(%s, %s) got:%v, expected:%v", v.str1, v.str2, got, v.want)
		}
	}
}
