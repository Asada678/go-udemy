package main

import (
	"fmt"
	"strconv"
)

// error handling

func main() {
	var s string = "100a"

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i)
}
