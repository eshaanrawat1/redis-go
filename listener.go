package main

import (
	"bytes"
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		recvBuf := make([]byte, 1024)

		_, err := conn.Read(recvBuf)
		if err != nil {
			fmt.Println("Error reading from TCP stream: ", err)
		}

		recvBuf = bytes.Trim(recvBuf, "\x00")

		res := string(recvBuf)
		res = strings.Replace(res, "\n", "", -1)

		if res == "Quit" {
			fmt.Println("Quitting ...")
			return
		}

		if res == "hello" {
			fmt.Println("hi")
		} else {
			fmt.Println(res)
		}
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
