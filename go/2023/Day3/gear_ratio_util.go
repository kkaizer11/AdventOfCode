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

func line_len() (int, int) {
	txt := read()
	size := 0
	lines := 0
	for is_line_break(txt[size]) == false {
		size++
	}
	for _, chr := range txt {
		if is_line_break(uint8(chr)) == true {
			lines++
		}
	}
	return size - 1, lines - 1
}
