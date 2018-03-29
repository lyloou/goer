package slice

import "testing"

var s = []int{0, 1, 2, 3, 4, 5}
func TestRotate(t *testing.T) {
	t.Log(s)
	Rotate(s, 2)
	t.Log(s)
}
