package main

import "fmt"

// type switch

func anything(a interface{}) {
	// fmt.Println(a)
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 10000)
	}
}

func main() {
	anything("aaa")
	anything(1)

	var x interface{} = "3"
	i, isInt := x.(int) // 型アサーション
	fmt.Println(i + 2)
	fmt.Println(i, isInt)
	// fmt.Println(x + 2)

	f, isFloat64 := x.(float64) // 型アサーション
	fmt.Println(f, isFloat64)   // panic

	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("I don't know")
	}

	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't know")
	}

	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	}

}
