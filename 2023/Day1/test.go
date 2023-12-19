package main

import "fmt"

func main() {
	x := "banana"
	y := x[0:3]
	fmt.Println(len(x))
	fmt.Printf("Tipo: %T", y[1])
}
