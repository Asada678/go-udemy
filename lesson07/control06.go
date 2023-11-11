package main

import "fmt"

// label for

func main() {
Loop:
	for {
		for {
			for {
				fmt.Println("START")
				break Loop
			}
			fmt.Println("do nothing 1")
		}
		fmt.Println("do nothing 2")
	}
	fmt.Println("END")

Loop2:
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if j > 1 {
				continue Loop2
			}
			fmt.Println(i, j, i*j)
		}
		fmt.Println("do nothing")
	}
}
