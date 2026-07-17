package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		recvBuf := make([]byte, 1024)

		_, err := conn.Read(recvBuf)
		if err != nil {
			fmt.Println("Error reading from TCP stream: ", err)
		}

		fmt.Println(recvBuf)
	}

}

func Listen() {
	ln, err := net.Listen("tcp", ":5284")
	if err != nil {
		fmt.Println("Error happened: ", err)
	}
	fmt.Println("Server is listening...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error in accepting connection: ", err)
		}

		fmt.Println("Accepted a TCP connection")
		go handleConnection(conn)
	}
}
