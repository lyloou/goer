package array

func Filter(arr []interface{}, f func(a interface{}) bool) []interface{} {
	for i := len(arr) - 1; i >= 0; i-- {
		if !f(arr[i]) {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

func FilterFromInt64Array(arr []int64, f func(a interface{}) bool) []int64 {
	return ToInt64Array(Filter(ConvertInt64ArrayTo(arr), f))
}

func FilterFromStringArray(arr []string, f func(a interface{}) bool) []string {
	return ToStringArray(Filter(ConvertStringArrayTo(arr), f))
}
