package main

import (
	"log"
	"os"
)

func read() string {
	text, err := os.ReadFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(text)
}

func is_line_break(c uint8) bool {
	if c == '\n' {
		return true
	}
	return false
}
