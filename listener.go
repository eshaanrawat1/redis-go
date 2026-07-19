package main

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

var cache map[string]int = make(map[string]int)

func handleGet(key string) (int, error) {
	value, ok := cache[key]
	if !ok {
		return -1, errors.New("Key not found in cache")
	}
	return value, nil
}

func handleSet(key string, value int) {
	cache[key] = value
}

func handleDelete(key string) {
	_, ok := cache[key]
	if ok {
		delete(cache, key)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		recvBuf := make([]byte, 1024)

		_, err := conn.Read(recvBuf)
		if err != nil {
			fmt.Println("Error reading from TCP stream: ", err)
			return
		}

		recvBuf = bytes.Trim(recvBuf, "\x00")

		res := string(recvBuf)
		res = strings.Replace(res, "\n", "", -1)

		cmds := strings.Split(res, " ")

		if cmds[0] == "Exit" {
			fmt.Println("Exiting ...")
			return
		}

		switch cmds[0] {
		case "Exit":
			fmt.Println("Exiting ...")
			return
		case "GET":
			value, err := handleGet(cmds[1])
			if err != nil {
				fmt.Println("Error parsing arguments for Get")
				return
			}
			fmt.Println(cmds[1], ": ", value)
		case "SET":
			num, err := strconv.Atoi(cmds[2])
			if err != nil {
				fmt.Println("Error parsing arguments for Set")
				return
			}
			handleSet(cmds[1], num)
			fmt.Println("Set: ", cmds[1], "to: ", num)
		case "DELETE":
			handleDelete(cmds[1])
			fmt.Println("Deleted: ", cmds[1])
		default:
			fmt.Println("Invalid command detected: ", cmds[0])
		}
		fmt.Println(cache)
	}

}

func Listen() {
	ln, err := net.Listen("tcp", ":5284")
	if err != nil {
		fmt.Println("Error happened: ", err)
		return
	}
	fmt.Println("Server is listening...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error in accepting connection: ", err)
			return
		}

		fmt.Println("Accepted a TCP connection")
		go handleConnection(conn)
	}
}
