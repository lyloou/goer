package main

import (
	"github.com/lyloou/goer/cmd/project-euler/handler"
	"bufio"
	"fmt"
	"strings"
	"sort"
	"github.com/lyloou/goer/pkg/array"
	"io/ioutil"
	"strconv"
)

// http://project-euler-answers-in-go.blogspot.com/2011/07/problem-22.html

func main() {
	// read
	ss := readFromFile()

	//sort
	sort.Strings(ss)

	total := 0
	for k, v := range ss {
		total += (k + 1) * getSum(v)
	}
	fmt.Println(total)
}

func readFromFile() []string {
	file, err := ioutil.ReadFile("022-names-scores.txt")
	if err != nil {
		panic(err)
	}
	s := string(file)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "\"", "", -1)
	return strings.Split(s, ",")
}

func readFromFile2() []string {
	s := ""
	handler.HandleScannerWithFilename("022-names-scores.txt", func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			s += scanner.Text()
		}
		return nil
	})

	// https://stackoverflow.com/questions/40760490/how-to-trim-char-from-a-string-in-golang
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "\"", "", -1)
	return strings.Split(s, ",")
}

func getSum(s string) int {
	sum := 0
	for _, v := range s {
		sum += int(v) - int('A') + 1
	}
	return sum
}

func getSum2(s string) int {
	sum := 0
	for _, v := range s {
		i, err := strconv.Atoi(fmt.Sprintf("%d", v-'A'))
		if err != nil {
			panic(err)
		}
		sum += i + 1
	}
	return sum
}

func getIndex(ss []string, s string) int {
	return array.Index(array.ConvertStringArrayTo(ss), s, func(item1, item2 interface{}) bool {
		return item1 == item2
	}) + 1
}
