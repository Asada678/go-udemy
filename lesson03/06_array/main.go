package main

import "fmt"

// array

func main() {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1) // -> [3]int :配列は要素数を含めて型として定義される

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"C", "D"} // ...は要素数を自動でカウントする
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4) // -> [2]string

	fmt.Println(arr1[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2-1])

	arr2[2] = "C"
	fmt.Println(arr2)

	// var arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5)

	fmt.Println(len(arr1))
}
