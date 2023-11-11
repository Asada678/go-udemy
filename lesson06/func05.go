package main

import "fmt"

// closure

func Later() func(string) string {
	var store string
	return func(next string) string {
		{
			s := store
			store = next
			return s
		}
	}
}

func main() {
	f := Later()
	fmt.Println(f("Hello"))
	fmt.Println(f("my"))
	fmt.Println(f("name"))
	fmt.Println(f("is"))
	fmt.Println(f("Golang"))
}
