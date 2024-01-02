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

func line_len() int {
	txt := read()
	size := 0
	for is_line_break(txt[size]) == false {
		size++
	}
	return size - 1
}
