package method

import (
	"fmt"
	"strings"
	"time"
)

func ExamplePrint() {
	Print(time.Hour)
	fmt.Println()
	Print(new(strings.Replacer))

	// Output:
	//type time.Durationfunc (time.Duration) Hours() float64
	//func (time.Duration) Minutes() float64
	//func (time.Duration) Nanoseconds() int64
	//func (time.Duration) Round(time.Duration) time.Duration
	//func (time.Duration) Seconds() float64
	//func (time.Duration) String() string
	//func (time.Duration) Truncate(time.Duration) time.Duration
	//
	//type *strings.Replacerfunc (*strings.Replacer) Replace(string) string
	//func (*strings.Replacer) WriteString(io.Writer, string) (int, error)
}
