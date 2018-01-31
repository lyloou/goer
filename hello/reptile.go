package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	resp, err := http.Get("https://baidu.com")
	if err != nil {
		//fmt.Println("http get error!")
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error!")
		return
	}

	src := string(body)
	//fmt.Println(src)

	//fmt.Println("=============================><=========================")
	// 转换标签为小写
	re, _ := regexp.Compile("<[\\S\\s]+?>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//fmt.Println(src)

	// 去除 style 标签及其内容
	//re, _ = regexp.Compile("<style[\\S\\s]+?</style>")
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//fmt.Println(src)

	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//fmt.Println(src)

	re, _ = regexp.Compile("\\<div[\\S\\s]+?\\</div\\>")
	src = re.ReplaceAllString(src, "")
	//fmt.Println(src)

	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "")
	//fmt.Println(src)

	regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "")
	//fmt.Println(src)

	fmt.Println(strings.TrimSpace(src))
	var dk = "      ksdf dslkfj  sdkfj          dkfjke                        " +
		"sdkfj               "
	println(strings.TrimSpace(dk))
}
