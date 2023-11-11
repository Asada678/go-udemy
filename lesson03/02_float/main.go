package main

import "fmt"

// float

func main() {
	/*
		// float32 is the set of all IEEE-754 32-bit floating-point numbers.
		type float32 float32

		// float64 is the set of all IEEE-754 64-bit floating-point numbers.
		type float64 float64
	*/
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2 // 暗黙的でもfl64となる
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	fmt.Printf("%T\n", float64(fl32))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)
}
