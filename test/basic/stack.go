package basic

import "fmt"

// https://yar999.gitbooks.io/gopl-zh/content/ch7/ch7-14.html
type Stack []string

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

var EmptyStackError = fmt.Errorf("stack is empty")

func (s *Stack) Pop() (string, error) {
	if len(*s) == 0 {
		return "", EmptyStackError
	}
	result := (*s)[(len(*s) - 1)]
	*s = (*s)[:(len(*s) - 1)]
	return result, nil
}
