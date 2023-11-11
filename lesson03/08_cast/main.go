package main

import (
	"fmt"
	"strconv"
)

// cast

func main() {
	var i int = 1
	fl64 := float64(i)
	fmt.Println(fl64)
	fmt.Printf("i = %T\n", i)
	fmt.Printf("fl64 = %T\n", fl64)

	i2 := int(fl64)
	fmt.Printf("i2 = %T\n", i2)

	fl := 10.5
	i3 := int(fl) // 小数点以下切り捨て
	fmt.Printf("i3 = %T\n", i3)
	fmt.Println(i3)

	var s string = "100a"
	fmt.Printf("s = %T\n", s)

	i4, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i4 = %T\n", i4)
	fmt.Println(i4)
	
	var i5 int = 500
	s2 := strconv.Itoa(i5)
	fmt.Printf("s2 = %T\n", s2)
	fmt.Println(s2)

	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)
}
