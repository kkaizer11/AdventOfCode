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

func sum_digits(s string) int {
	digits := map[uint8]int{}
	for i := 0; i <= 9; i++ {
		digits['0'+uint8(i)] = i
	}
	fmt.Println(digits)
	fmt.Println(digits['8'])
	return 0
}

func main() {
	sum_digits("banan")
}
