package array

func All(arr []interface{}, f func(interface{}) bool) bool {
	for _, v := range arr {
		if !f(v) {
			return false
		}
	}
	return true
}
