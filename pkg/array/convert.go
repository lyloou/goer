package array

func ConvertInt64ArrayTo(arr []int64) []interface{} {
	s := make([]interface{}, len(arr))
	for i, v := range arr {
		s[i] = v
	}
	return s
}

func ConvertStringArrayTo(arr []string) []interface{} {
	s := make([]interface{}, len(arr))
	for i, v := range arr {
		s[i] = v
	}
	return s
}
