package crypto

import (
	"crypto/rand"
	"fmt"
	"testing"
	"hash/crc32"
	"io/ioutil"
	"crypto/sha1"
)

func TestCrypto4(t *testing.T) {
	file, err := ioutil.ReadFile("text2.txt") // by change file to see their different
	if err  != nil {
		fmt.Println(err)
		return
	}

	h := sha1.New()

	h.Write(file)
	bs := h.Sum([]byte{})
	fmt.Println(bs)

}
// https://www.golang-book.com/books/intro/13#section6
func TestCrypto3(t *testing.T) {
	h, err := getHash("text1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	h2, err := getHash("text2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(h, h2, h == h2)
}
func getHash(filename string) (uint32, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	h := crc32.NewIEEE()
	h.Write(file)
	return h.Sum32(), nil
}

func TestCrypto2(t *testing.T) {
	h := crc32.NewIEEE()
	h.Write([]byte("test"))
	v := h.Sum32()
	fmt.Println(v)
}

func TestCrypto1(t *testing.T) {
	fmt.Println("vim-go")
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%x", buf))
}
