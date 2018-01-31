package main

import (
	log "github.com/sirupsen/logrus"
	"fmt"
)

func main() {
	fmt.Println("hi hello")
	log.WithFields(log.Fields{
		"animal":  []byte{1, 2, 3},
		"animal1": func() { fmt.Println("hi") },
		"animal2": "walrus2",
		"animal3": "walrus3",
	}).Warn("A walrus appears")
}
