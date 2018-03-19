package util

import (
	"fmt"
	"testing"
)

const BASE_URL = "http://localhost:8080"

func getData(url, path string) []byte {
	data, err := HttpGet(fmt.Sprintf("%v%v", url, path))
	if err != nil {
		panic(err)
	}
	return data
}

func postData(url, path string, param []byte) []byte {
	data, err := HttpPost(fmt.Sprintf("%v%v", url, path), param)
	if err != nil {
		panic(err)
	}
	return data
}

func TestHttpGet(t *testing.T) {
	fmt.Println(string(getData(BASE_URL, "/api/get")))
}

func TestHttpPost(t *testing.T) {
	fmt.Println(string(postData(BASE_URL, "/api/post", []byte("[55,43,332,333]"))))
}
