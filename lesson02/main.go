package main

import "fmt"

// variable

// i5 := 500
var i5 int = 500

func outer() {
	s4 := "outer s4"
	fmt.Println(s4)

	var s5 string = "outer s5"
	fmt.Println(s5)
}

func main() {
	// 明示的な定義
	// var variable type = value
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int    // 0
	var s3 string // ""
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	// 暗黙的な定義
	// variable := value
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// i4 := 500
	// i4 = "Hello"
	// fmt.Println(i4)

	fmt.Println(i5)

	outer()
}
