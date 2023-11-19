package main

import "fmt"

// slice
// variable length argument

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}

	return n
}

func main() {
	fmt.Println(Sum(1, 3, 5))
	fmt.Println(Sum(1, 3, 5, 7, 9, 11, 13))
	fmt.Println(Sum())

	sl := []int{1, 2, 3}
	fmt.Println(Sum(sl...))
}
