package main

import "fmt"

// int

func main() {
	var i int = 100

	/*
		intは明示的に宣言しなければ、環境によって、32, 64 が割り当てられる

		// int8 is the set of all signed 8-bit integers.
		// Range: -128 through 127.
		type int8 int8

		// int16 is the set of all signed 16-bit integers.
		// Range: -32768 through 32767.
		type int16 int16

		// int32 is the set of all signed 32-bit integers.
		// Range: -2147483648 through 2147483647.
		type int32 int32

		// int64 is the set of all signed 64-bit integers.
		// Range: -9223372036854775808 through 9223372036854775807.
		type int64 int64
	*/

	var i2 int64 = 200

	fmt.Println(i + 50)
	fmt.Println(i2 + 50)

	fmt.Printf("%T\n", i2)

	fmt.Printf("%T\n", int32(i2))

	fmt.Println(i + int(i2))

}
