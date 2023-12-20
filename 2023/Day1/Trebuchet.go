package main

import (
	"log"
	"os"
)

func read() string {
	text, err := os.ReadFile("doc.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(text)
}

func str_is_digit(c uint8) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func is_line_break(c uint8) bool {
	if c == '\n' {
		return true
	}
	return false
}
