package method

import "sync"

// https://yar999.gitbooks.io/gopl-zh/content/ch6/ch6-03.html
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) (string, bool) {
	cache.Lock()
	defer cache.Unlock()
	v,ok := cache.mapping[key]
	return v,ok
}

func Insert(key, value string) {
	cache.Lock()
	defer cache.Unlock()
	cache.mapping[key] = value
}
