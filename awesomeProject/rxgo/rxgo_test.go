package rxgo

import (
	"testing"
	"fmt"
	"github.com/reactivex/rxgo/observer"
	"github.com/reactivex/rxgo/iterable"
	"github.com/reactivex/rxgo/observable"
	"errors"
	"github.com/reactivex/rxgo/handlers"
	"time"
)

func TestRxgo(t *testing.T) {
	f1 := func() interface{} {
		time.Sleep(2 * time.Second)
		return 1
	}

	f2 := func() interface{} {
		time.Sleep(2 * time.Second)
		return 2
	}

	onNext := handlers.NextFunc(func(v interface{}) {
		val := v
		fmt.Printf("%v, %p, %p\n", v, &val, &v)
	})
	sub := observable.Start(f1, f2).Subscribe(onNext)
	result := <-sub
	if err := result.Err(); err != nil {
		fmt.Println("出错了啦")
	}

}
func sample2() {
	score := 9
	onNext := handlers.NextFunc(func(item interface{}) {
		if num, ok := item.(int); ok {
			fmt.Println("next:", score)
			score += num
		}
	})
	onDone := handlers.DoneFunc(func() {
		fmt.Println("done:", score)
		score *= 2
	})
	watcher := observer.New(onNext, onDone)
	sub := observable.Just(1).Subscribe(watcher)
	fmt.Println(score)
	<-sub
	fmt.Println(score)
}
func sample1() {
	watcher := observer.Observer{
		NextHandler: func(item interface{}) {
			fmt.Printf("NextHandler: %v\n", item)
		},
		ErrHandler: func(err error) {
			fmt.Printf("ErrHandler: %v\n", err)
		},
		DoneHandler: func() {
			fmt.Println("DoneHandler")
		},
	}
	it, _ := iterable.New([]interface{}{1, 2, 3, 4, errors.New("bang"), 5})
	source := observable.From(it)
	sub := source.Subscribe(watcher)
	<-sub
}
