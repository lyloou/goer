package intset_test

import (
	"fmt"
	"testing"

	"github.com/lyloou/goer/test/basic/intset"
)

func TestIntSet(t *testing.T) {
	var is intset.IntSet

	is.Add(129)
	//is.Add(129)
	is.Add(331)
	is.Add(1800)
	fmt.Println(&is)
	fmt.Println(is.String())
	fmt.Println(is)
}
