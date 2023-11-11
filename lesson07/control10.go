package main

import "fmt"

// init

func init() {
	fmt.Println("init")
}

// 出現順に実行される
func init() {
	fmt.Println("init2")
}

func main() {
	fmt.Println("Main")
}
