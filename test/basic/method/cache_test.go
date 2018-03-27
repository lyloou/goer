package method_test

import (
	"testing"

	"github.com/lyloou/goer/test/basic/method"
)

func TestLookup(t *testing.T) {
	method.Insert("a", "aaa")
	method.Insert("b", "aab")
	method.Insert("c", "aac")
	method.Insert("d", "aad")

	t.Log(method.Lookup("a"))
	t.Log(method.Lookup("b"))
	t.Log(method.Lookup("c"))
	t.Log(method.Lookup("d"))
	t.Log(method.Lookup("e"))
}
