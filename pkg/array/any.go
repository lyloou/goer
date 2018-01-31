package array

func Any(arr []interface{}, f func(interface{}) bool) bool {
	for _, v := range arr {
		if f(v) {
			return true
		}
	}
	return false
}
