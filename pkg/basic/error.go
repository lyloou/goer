package basic

import "strings"

// IsSameError determine whether the two errors are the same
func IsSameError(e1 error, e2 error) bool {
	if e1 == nil && e2 == nil {
		return true
	}
	if e1 != nil && e2 != nil && strings.EqualFold(e1.Error(), e2.Error()) {
		return true
	}
	return false
}
