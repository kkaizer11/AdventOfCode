package main

import (
	"fmt"
	"log"
	"os"
)

func calibration_value(text []uint8) int {
	keyval := map[int]int{}
	for i := 0; i < 9; i++ {
		keyval[48+i] = i
	}
	keyval[10] = -1
}

func main() {
	txt, err := os.ReadFile("doc.txt")
	x := []rune("1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(txt)
	fmt.Printf("%T, %d\n", x, x)
}
