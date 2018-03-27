package counter

import (
	"fmt"
	"testing"
)

func TestCounter(t *testing.T) {
	var c Counter
	for i := 0; i < 10; i++ {
		c.Increment()
		if c.N() == 8 {
			c.Reset()
		}
		fmt.Println(i, c.N())
	}
	fmt.Println(c.N())
}
