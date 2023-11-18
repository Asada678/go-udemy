package main

import "fmt"

// slice
// copy

func main() {
	sl := []int{100, 200}
	sl2 := sl

	sl2[0] = 1000
	fmt.Println(sl) // 参照型はコピー元も変更される

	var i int = 10
	i2 := i
	i2 = 100
	fmt.Println(i, i2) // 基本型はそれぞれ別

	sl3 := []int{1, 2, 3, 4, 5}
	sl4 := make([]int, 5, 10)
	fmt.Println(sl4)
	n := copy(sl4, sl3)
	fmt.Println(n, sl4)
}
