package array

func Map(arr []interface{}, f func(interface{}) interface{}) []interface{} {
	for i, v := range arr {
		arr[i] = f(v)
	}
	return arr
}
