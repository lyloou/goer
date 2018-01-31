package array

func Include(arr []interface{}, item interface{}, isEqual func(item1, item2 interface{}) bool) bool {
	return Index(arr, item, isEqual) > 0
}
