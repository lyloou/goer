
```go
var buf *bytes.Buffer
buf = new(bytes.Buffer) // buf 变量赋予了一个*bytes.Buffer的空指针
fmt.Println(buf == nil) // false。 buf的动态类型和动态值必须同时为nil，该结果才为true;
                        // https://yar999.gitbooks.io/gopl-zh/content/ch7/ch7-05.html
```

```go
log.Fatal("hini") // Fatal is equivalent to Print() followed by a call to os.Exit(1).
```


## 类型转换
```go
package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {

	var w io.Writer
	w = new(bytes.Buffer)

	w.Write([]byte("hello"))
	if v, ok := w.(io.Writer); ok {
		if vv, ok := w.(*bytes.Buffer); ok {
			fmt.Println("if buffer:", vv.String())
		}
		fmt.Println("if writer:", v)
	}

	switch v := w.(type) {
	case *bytes.Buffer:
		fmt.Println("switch buffer:", v)
	case io.Writer:
		fmt.Println("switch writer:", v)
	}
}

// Output:
//if buffer: hello
//if writer: hello
//switch buffer: hello

```