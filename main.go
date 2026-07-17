package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parser() {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		cmds := strings.Split(text, " ")

		if cmds[0] == "Quit" {
			break
		}

		if cmds[0] == "Hi" {
			fmt.Println("hello there")
		}
	}
}

func main() {
	parser()
}
