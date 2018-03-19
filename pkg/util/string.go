package util

// Substr returns the substr from start to length.
func Substr(s string, start, length int) string {
	bt := []rune(s)
	if start < 0 {
		start = 0
	}
	if start > len(bt) {
		start = start % len(bt)
	}
	var end int
	if (start + length) > (len(bt) - 1) {
		end = len(bt)
	} else {
		end = start + length
	}
	return string(bt[start:end])
}

func IsContainsInInt64Slice(arr []int64, item int64) bool{
	for _, v := range arr{
		if item == v {
			return true
		}
	}
	return false
}
func IsContainsInIntSlice(arr []int, item int) bool{
	for _, v := range arr{
		if item == v {
			return true
		}
	}
	return false
}
