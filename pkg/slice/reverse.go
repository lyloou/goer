package slice

// https://yar999.gitbooks.io/gopl-zh/content/ch4/ch4-02.html
// reverse reverses a slice of ints in place.
func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Rotate(s []int, n int) {
	Reverse(s[:n])
	Reverse(s[n:])
	Reverse(s[:])
}
