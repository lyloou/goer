package main

import (
	"os"
	"fmt"
	"regexp"
)

func main() {
	a := "I am learning Go language"
	matched, err := regexp.MatchString("[a-z]{2,4}", a)
	re, _ := regexp.Compile("[a-z]{2,4}")
	fmt.Println(matched, err)
	allString := re.ReplaceAllString(a, "oo")
	fmt.Println(allString)
	allFunc := re.ReplaceAllFunc([]byte(a), func(bytes []byte) []byte {
		str := string(bytes)
		str += "~"
		return []byte(str)
	})
	fmt.Println(string(allFunc))

	string2 := re.ExpandString([]byte(a), "222", a, []int{1,2})
	fmt.Println(string(string2))

}
func sample2(a string) {
	re, _ := regexp.Compile("[a-z]{2,4}")
	one := re.Find([]byte(a))
	fmt.Println(string(one))
	all := re.FindAll([]byte(a), -1)
	for k, v := range all {
		fmt.Println("k=", k, " v=", string(v))
	}
	fmt.Println()
	allIndex := re.FindAllIndex([]byte(a), -1)
	fmt.Println(allIndex)
	re2, _ := regexp.Compile("am(.*)lang(.*)")
	fs := re2.FindSubmatch([]byte(a))
	for _, v := range fs {
		fmt.Print(string(v), ",")
	}
	fmt.Println()
	submatch := re2.FindAllSubmatch([]byte(a), -1)
	for _, v := range submatch {
		for _, vv := range v {
			fmt.Print(string(vv))
		}
	}
	fmt.Println("==============")
	submatchindex := re2.FindAllSubmatchIndex([]byte(a), -1)
	for _, v := range submatchindex {
		fmt.Print((v))
	}
}

func sample1() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: regexp [string]")
		os.Exit(1)
	} else if m, _ := regexp.MatchString("[0-9]+", os.Args[1]); m {
		fmt.Println("数字")
	} else {
		fmt.Println("不是数字")
	}
}
