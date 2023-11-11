package main

import "fmt"

// panic recover

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("runtime error 123")
	fmt.Println("Start")
}
