package main

import "fmt"

func main() {
	txt := read()
	txt_lines := gnl(txt)
	fmt.Println(txt_lines)
	fmt.Println(len(txt_lines))
}
