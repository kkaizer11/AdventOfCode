package main

import (
	"fmt"
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

func main() {
	fmt.Println(str_is_digit('N'))
	fmt.Println(str_is_digit('0'))
	fmt.Println(str_is_digit('9'))
	fmt.Println(str_is_digit('5'))
}
