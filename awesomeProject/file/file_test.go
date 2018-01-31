package file

import (
	"testing"
	"os"
	"fmt"
)

func TestFile(t *testing.T) {

}


func sample2() {
	fName := "lou/hi.txt"
	f, err := os.Create(fName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	m, _ := f.WriteString("我们的世界是多么的美丽啊！")
	n, _ := f.Write([]byte("你好， 你知道了吗？"))
	f.WriteAt([]byte("对不起你错了！"), int64(m+n))
	f.Close()

	// 重新打开
	f, err = os.Open(fName)
	defer f.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
func sample1() {
	os.MkdirAll("lou/lou", 0777)
	os.MkdirAll("lou/lou2", 0777)
	os.MkdirAll("lou/lou23", 0777)
	os.MkdirAll("lou/lou23/ad", 0777)
	os.MkdirAll("lou/lou23/ab", 0777)
	os.RemoveAll("lou/lou23")
	os.Create("lou/hi.txt")
	os.NewFile(777, "lou/newhi.txt")
	os.Open("lou/hi.txt")
	os.OpenFile("lou/hi.txt", 0777, 0777)
}
