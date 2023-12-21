package main

import "fmt"

func main() {
	txt := read()
	txt_lines := gnl(txt)
	sum := sum_all(txt_lines)
	fmt.Println(sum)
}
