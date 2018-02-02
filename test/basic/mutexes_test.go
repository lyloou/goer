package basic

import (
	"testing"
	"fmt"
	"time"
	"sync"
	"errors"
)

func TestMutexes(t *testing.T) {
	var count int64
	var total int64 = 10
	c := make(chan int64)
	m := new(sync.Mutex)
	for i := int64(0); i < total; i++ {
		go func(i int64) {
			m.Lock()
			c <- i
			fmt.Println()
			fmt.Println(i, "start")

			if count > total {
				fmt.Println("已经结束了")
				m.Unlock()
				return
			}

			result, err := divide(total, i)
			time.Sleep(time.Second)

			if err != nil {
				fmt.Println(i, "no - end")
				fmt.Println(err)
			} else {
				fmt.Println(i, ":", total, "/", i, "=", result)
				fmt.Println(i, "end")
			}
			count++
			if count == total {
				close(c)
			}

			m.Unlock()
		}(i)
	}

	for v := range c {
		fmt.Println(v, "main")
	}
}

func divide(first, second int64) (int64, error) {
	if second == 0 {
		return 0, errors.New("除数为0")
	}
	return first / second, nil
}
