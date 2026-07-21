package main

import (
	"bufio"
	"os"
)

type Logger struct {
	Path string
}

func NewLogger() Logger {
	return Logger{
		Path: "./logs.txt",
	}
}

func (l *Logger) WriteLog(msg string) {
	file, err := os.OpenFile(l.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	if _, err := file.WriteString(msg); err != nil {
		panic(err)
	}
}

func (l *Logger) ReadLog() []string {
	file, err := os.OpenFile(l.Path, os.O_RDONLY, 0444)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	cmds := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		cmds = append(cmds, line)
	}
	return cmds
}
