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

func str_is_digit(c rune) bool {
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

func two_digits(s string) string {
	num_str := ""
	for _, word := range s {
		if str_is_digit(word) == true {
			num_str += string(word)
		}
	}
	// TODO: funçao está bugada
	if len(num_str) >= 1 {
		return num_str
	}
	return string(num_str[0] + num_str[len(num_str)-1])
}

func gnl(s string) []string {
	var (
		bp     int = 0
		result []string
	)
	for i := 0; i < len(s); i++ {
		if is_line_break(s[i]) {
			result = append(result, two_digits(s[bp:i]))
			bp = i + 1
		}
	}
	return result
}
