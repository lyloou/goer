package main

import (
	"fmt"
	"io"
	"crypto/md5"
	"crypto/aes"
	"crypto/cipher"
	"golang.org/x/crypto/scrypt"
)

var commonIv = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func main() {
	plaintext := []byte("my name is Astaxie")
	key_text := "astaxie12798akljzmknm.ahkjkljl;k"

	fmt.Println(len(key_text))

	c, _ := aes.NewCipher([]byte(key_text))

	cfb := cipher.NewCFBEncrypter(c, commonIv)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("%s=>%x", plaintext, ciphertext)
	fmt.Println()

	cfbdec := cipher.NewCFBDecrypter(c, commonIv)
	plaintextCopy := make([]byte, len(plaintext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	fmt.Printf("%x=>%s", ciphertext, plaintextCopy)
}

func sample_1() {
	// h := sha256.New()
	h := md5.New()
	io.WriteString(h, "hi hello")
	fmt.Printf("% x", h.Sum(nil))
	fmt.Println()
	key, _ := scrypt.Key([]byte("somepassword"), []byte("@#$%"), 16384, 8, 1, 32)
	fmt.Printf("%x", key)
}
