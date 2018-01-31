package main

import (
	"os"
	"fmt"
	"net"
)


func main() {
	ipAddr, err := net.ResolveIPAddr("tcp", "192.168.0.51:80")
	if err != nil {
		fmt.Println(err, ipAddr)
		return
	}
	fmt.Println(ipAddr)
}
func sample1() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	ip := net.ParseIP(name)
	if ip == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", ip.String())
	}
	os.Exit(0)
}
