package main

import "fmt"

func main() {
	txt := read()
	fmt.Println(txt)
	size := 0
	for i := 0; is_line_break(txt[i]) == false; i++ {
		size++
	}
	fmt.Println(size)
	fmt.Println(txt[size])
}
