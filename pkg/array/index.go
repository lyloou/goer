package array

func Index(arr []interface{}, item interface{}, isEqual func(item1, item2 interface{}) bool) int {
	for i, v := range arr {

		if isEqual(v, item) {
			return i
		}
	}
	return -1
}
