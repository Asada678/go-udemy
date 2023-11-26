package main

import (
	"log"
	"os"
)

// os

func main() {

	// Exit
	/*
		os.Exit(1)
		fmt.Println("Start")
	*/

	/*

		defer func() {
			fmt.Println("defer")
			}()
			os.Exit(0) // status0で終了した場合はdefer文が実行されない
	*/

	_, err := os.Open("A.txt")
	if err != nil {
		log.Fatalln(err)
	}

}
