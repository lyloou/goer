package math_test

import (
	"errors"
	"testing"

	"github.com/lyloou/goer/pkg/basic"
	"github.com/lyloou/goer/pkg/math"
)

// https://yar999.gitbooks.io/gopl-zh/content/ch11/ch11-02.html
// 表格驱动测试
func TestDivide(t *testing.T) {
	var tests = []struct {
		a   int64
		b   int64
		r   int64
		err error
	}{
		{4, 2, 2, nil},
		{4, 3, 1, nil},
		{2, 4, 0, nil},
		{4, 0, 0, errors.New("division by zero")},
	}

	for _, test := range tests {
		if r, err := math.Divide(test.a, test.b); r != test.r || !basic.IsSameError(err, test.err) {
			t.Errorf("Divide(%d, %d) got (%d, %v), want (%d, %v)", test.a, test.b, r, err, test.r, test.err)
		}
	}
}

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Divide(3009, 107)
	}
}
