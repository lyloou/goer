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
