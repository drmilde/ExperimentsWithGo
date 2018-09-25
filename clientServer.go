package main

import (
	"fmt"
	"net"
)

func server() {
	ln, err := net.Listen("tcp", ":9999")

	if err != nil {
		fmt.Println("fehler im server")
		return
	}

	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleServerConnection(c)
	}

}

func handleServerConnection(c net.Conn) {
	// behandele die Server Connection
}

// starte client server
func main() {

}
