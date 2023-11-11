package main

import "fmt"

// arg func

func CallFunction(f func()) {
	f()
}

func main() {
	CallFunction(func() {
		fmt.Println("I'm a function")
	})
}
