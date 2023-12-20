package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	tx := "pqr3stu8vwx"
	tx = two_digits(tx)
	tx_int, err := strconv.Atoi(tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx_int)
}
