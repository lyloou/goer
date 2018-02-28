package timing

import (
	"fmt"
	"os"
	"time"
)

func Start(a ...interface{}) (start time.Time) {
	fmt.Fprintln(os.Stdout, a...)
	return time.Now()
}

func Stop(start time.Time, a ...interface{}) (spending time.Duration) {
	fmt.Fprintln(os.Stdout, a...)
	return time.Now().Sub(start)
}
