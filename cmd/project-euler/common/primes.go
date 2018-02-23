package common

import "math"

func IsPrime(n int64) bool {
	s := int64(math.Sqrt(float64(n))) + 1
	for i := int64(2); i < s; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

