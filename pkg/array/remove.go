package array

func Remove(arr []interface{}, f func(a interface{}) bool) []interface{} {
	for i := len(arr) - 1; i >= 0; i-- {
		if f(arr[i]) {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

func RemoveFromInt64Array(arr []int64, f func(a interface{}) bool) []int64 {
	return ToInt64Array(Remove(ConvertInt64ArrayTo(arr), f))
}

func RemoveFromStringArray(arr []string, f func(a interface{}) bool) []string {
	return ToStringArray(Remove(ConvertStringArrayTo(arr), f))
}
