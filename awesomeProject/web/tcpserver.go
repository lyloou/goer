package main

import (
	"net"
	"fmt"
	"os"
	"time"
)

func main() {
	service := ":52214"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	fmt.Println(tcpAddr)
	listen, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		go handleConn(conn)
	}

}
func handleConn(conn net.Conn) {
	defer conn.Close()
	timeNow := time.Now().String()
	conn.Write([]byte(timeNow))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
