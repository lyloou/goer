package math

import "errors"

func Divide(x, y int64) (int64, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}
