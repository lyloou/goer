package common

import "testing"

func TestIsLeap(t *testing.T) {
	var tests = []struct {
		y      int
		isLeap bool
	}{
		{1900, false},
		{2003, false},

		{2000, true},
		{2004, true},
		{2008, true},
	}

	for _, tt := range tests {
		isLeap := IsLeap(tt.y)
		if isLeap != tt.isLeap {
			t.Errorf("when do year(%v) expect %v, got %v", tt.y, tt.isLeap, isLeap)
		}
	}
}

func BenchmarkIsPandigital(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPandigital("932718654")
	}
}

func BenchmarkIsPandigital2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPandigital2("932718654")
	}
}
