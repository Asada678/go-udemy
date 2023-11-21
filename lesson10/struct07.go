package main

import "fmt"

type MyInt int

func (mi MyInt) Print() {
	fmt.Println(mi)
}

func main() {
	var mi MyInt = 1
	fmt.Println(mi)
	fmt.Printf("%T\n", mi)

	i := 100
	fmt.Printf("%T\n", i)

	// sum := mi + i // 計算できない

	mi.Print()
}
