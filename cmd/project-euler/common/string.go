package common

import (
	"fmt"
	"sort"
	"strings"
)

func justShowOnce2(str string) bool {
	split := strings.Split(str, "")
	if len(split) < 2 {
		return true
	}
	m := make(map[string]string, 0)
	for _, v := range split {
		m[v] = v
	}
	return len(split) == len(m)
}

// https://github.com/polaris1119/The-Golang-Standard-Library-by-Example/blob/master/chapter02/02.1.md#212-子串出现次数字符串匹配
// JustShowOnce return true if 每个字符只显示一次。效果比上面的justShowOnce2要高。
func JustShowOnce(str string) bool {
	for _, v := range str {
		if strings.Count(str, fmt.Sprintf("%c", v)) > 1 {
			return false
		}
	}
	return true
}

func IsPermutation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	str1Sp := strings.Split(str1, "")
	str2Sp := strings.Split(str2, "")
	sort.Strings(str1Sp)
	sort.Strings(str2Sp)

	for k, v := range str1Sp {
		if str2Sp[k] != v {
			return false
		}
	}
	return true
}
