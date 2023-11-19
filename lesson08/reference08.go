package main

import "fmt"

// channel
// channelはキューの性質を持つFIFO
// 複数のゴルーチン感でのデータ授受を可能にするために設計されたデータ構造
// 宣言、操作

func main() {
	var ch1 chan int // 送受信可能

	// var ch2 <-chan int // 受信専用のチャネル

	// var ch3 chan<- int // 送信専用のチャネル

	ch1 = make(chan int)

	ch2 := make(chan int)

	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))

	ch3 <- 1
	fmt.Println(len(ch3))

	ch3 <- 10
	ch3 <- 100
	fmt.Println("len:", len(ch3))

	i := <-ch3
	fmt.Println(i) // -> 1 最初に送信した値
	fmt.Println("len:", len(ch3))

	i2 := <-ch3
	fmt.Println(i2)               // -> 10 2番目に送信した値
	fmt.Println("len:", len(ch3)) // 要素数が一つずつ減る

	fmt.Println(<-ch3)
	fmt.Println("len:", len(ch3))

	ch3 <- 1
	// fmt.Println("len:", len(ch3))
	// fmt.Println(<-ch3)
	// fmt.Println("len:", len(ch3))
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	ch3 <- 6 // バッファサイズ超過 -> deadlock!

}
