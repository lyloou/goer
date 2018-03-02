package common

import "math"

func GetDivisors(division int64) (divisors [] int64) {
	var i int64
	for i = 1; i <= division/2; i++ {
		if division%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return
}

func SumDivisors(divisors []int64) (sum int64) {
	for _, v := range divisors {
		sum += v
	}
	return
}

func IsPrime(n int64) bool {
	s := int64(math.Sqrt(float64(n))) + 1
	for i := int64(2); i < s; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func IsLeap(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	if year%4 == 0 {
		return true
	}
	return false
}

func Factorial(num int64) int64 {
	var result, i int64 = 1, 1
	for ; i <= num; i++ {
		result *= i
	}
	return result
}

func Gcd(num1, num2 int64) int64 {

	if num2 == 0 {
		return num1
	}
	return Gcd(num2, num1%num2)
}

// http://blog.csdn.net/u014609452/article/details/79115753
func FourPower(length int64) int64 {
	// 4^x = 2^(2x) = (2^x)^2
	var c int64 = 1 << (uint64(length) << 1)
	return c
}
