package main

import "fmt"

func calibration_value(s int) int {
	keyval := map[int]int{}
	for i := 0; i < 9; i++ {
		keyval[48+i] = i
	}
	keyval[10] = -1
	return keyval[s]
}

func main() {
	fmt.Println(calibration_value(54))
}
