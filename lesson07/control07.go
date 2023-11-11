package main

import (
	"fmt"
	"os"
)

// defer

func TestDefer() {
	defer fmt.Println("END") // 関数が終了時に実行される
	fmt.Println("START")
}

func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func main() {
	defer func() {
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()

	TestDefer()

	RunDefer()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.Write([]byte("Hello"))

}
