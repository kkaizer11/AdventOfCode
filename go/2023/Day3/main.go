package main

import "fmt"

// pouca coisa feita mas eu acredito
func main() {
	txt := read()
	line_size, line_cout := line_len()
	fmt.Println(txt)
	fmt.Println(line_cout)
	fmt.Println(line_size)
}
