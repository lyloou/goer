package main

import (
	"net"
	"fmt"
	"os"
	"time"
)

func main() {
	service := ":52214"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	fmt.Println(udpAddr)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	for {
		handleConn2(conn)
	}

}
func handleConn2(conn *net.UDPConn) {
	var buf[512]byte
	n, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}

	fmt.Println("========>",string(buf[0:n]))

	timeNow := time.Now().String()
	conn.WriteToUDP([]byte(timeNow+" hi hello"), addr)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
