package main

import (
	"bufio"
	"os"
	"strings"
)

func ReadFisrstLine() string {
	open, err := os.Open("log")
	defer open.Close()
	if err != nil {
		return ""
	}
	scanner := bufio.NewScanner(open)
	for scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func ProcessFirstLine() string {
	line := ReadFisrstLine()
	destline := strings.ReplaceAll(line, "11", "00")
	return destline
}
