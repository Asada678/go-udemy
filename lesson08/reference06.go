package main

import "fmt"

// map

func main() {
	var m = map[string]int{"A": 100, "B": 200}

	fmt.Println(m)

	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "Japan"
	m4[2] = "USA"
	fmt.Println(m4)

	fmt.Println(m["A"])
	fmt.Println(m4[2])
	fmt.Println(m4[3]) // 登録されていないキーに対しては基本型の初期値が返却される

	s, ok := m4[1]
	fmt.Println(s, ok)
	s2, ok2 := m4[10]
	fmt.Println(s2, ok2) // ok -> false
	if !ok2 {
		fmt.Println("error")
	}

	m4[2] = "US"
	m4[3] = "China"
	fmt.Println(m4)

	delete(m4, 2)
	fmt.Println(m4)

	fmt.Println(len(m4))

}
