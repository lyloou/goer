package basic

import (
	"testing"
)

type P int

func (p *P) Add(q P) {
	*p = *p + q
}

func (p P) Minus(q P) P {
	return p - q
}
func TestMethod(t *testing.T) {
	var p P
	q := &p
	q.Add(3)
	q.Add(2)
	q.Add(1)
	t.Log(*q)

	// 我们不能通过一个无法取到地址的接收器来调用指针方法，比如临时变量的内存地址就无法获取得到：
	// https://yar999.gitbooks.io/gopl-zh/content/ch6/ch6-02.html
	//P(2).Add(2) // cannot call pointer method on P(2)

	p = p.Minus(3)
	t.Log(p)
	p = P(3).Minus(12) // 非指针方法，可以这样用
	t.Log(p)
	p = (&p).Minus(32)
	t.Log(p)
}
