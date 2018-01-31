package array

func ToInt64Array(arr []interface{}) []int64 {
	if len(arr) < 1 {
		return []int64{}
	}

	if _, ok := arr[0].(int64); !ok {
		return []int64{}
	}

	s := make([]int64, len(arr))
	for i, v := range arr {
		s[i] = v.(int64)
	}
	return s
}

func ToStringArray(arr []interface{}) []string {
	if len(arr) < 1 {
		return []string{}
	}

	if _, ok := arr[0].(string); !ok {
		return []string{}
	}

	s := make([]string, len(arr))
	for i, v := range arr {
		s[i] = v.(string)
	}
	return s
}
