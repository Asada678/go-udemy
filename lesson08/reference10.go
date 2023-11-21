package main

import (
	"fmt"
	"time"
)

// channel
// close

func receiver(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")

}

func main() {
	ch1 := make(chan int, 2)

	ch1 <- 1

	close(ch1)

	// ch1 <- 1
	// fmt.Println(<-ch1)

	i, ok := <-ch1
	fmt.Println(i, ok)

	i2, ok := <-ch1
	fmt.Println(i2, ok)

	ch2 := make(chan int, 2)

	go receiver("1.goroutine", ch2)
	go receiver("2.goroutine", ch2)
	go receiver("3.goroutine", ch2)

	j := 0
	for j < 10 {
		ch2 <- j
		j++
	}
	close(ch2)
	time.Sleep(3 * time.Second)

}
