package basic

import (
	"testing"
	"fmt"
	"time"
	"math/rand"
)

// https://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly
func TestRandom(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(randomString(10))

	for i := 0; i < 10; i++ {
		fmt.Print(randInt(1, 10), " ")
	}
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
