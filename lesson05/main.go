package main

import "fmt"

// operator

func main() {
	fmt.Println(1 + 2)
	fmt.Println(5 - 1)
	fmt.Println(5 * 4)
	fmt.Println(61 / 3)
	fmt.Println(9 % 4)

	fmt.Println("ABC" + "DEF")

	n := 0
	n += 4
	fmt.Println(n)
	n++
	fmt.Println(n)
	n--
	fmt.Println(n)

	s := "ABC"
	s += "DEF"
	fmt.Println(s)

	fmt.Println(1 == 1)
	fmt.Println(1 == 2)
	fmt.Println(4 <= 8)
	fmt.Println(4 >= 8)
	fmt.Println(4 < 8)
	fmt.Println(4 > 8)

	fmt.Println(true == false)
	fmt.Println(true != false)

	fmt.Println(true && false == true)
	fmt.Println(true && false != true)

	fmt.Println(true || false == true)
	fmt.Println(false || false == true)

	fmt.Println(!true)
	fmt.Println(!false)
}
