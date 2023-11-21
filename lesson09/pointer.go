package main

import "fmt"

// pointer

func Double(i int) {
	i = i * 2
}

func DoubleV2(i *int) {
	*i = *i * 2
}

func DoubleV3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func main() {
	var i int = 100
	fmt.Println(i)

	fmt.Println(&i) // メモリ上のアドレス

	Double(i)
	fmt.Println(i) // -> 100 再出力しても結果は変わらない

	var p *int = &i // *でポインタ型を定義できる。&iでアドレスを渡す
	fmt.Println(p)
	fmt.Println(*p) // アドレスが指す実体を表示できる

	*p = 300
	fmt.Println(i)
	i = 200
	fmt.Println(*p)

	DoubleV2(&i)
	fmt.Println(i)

	DoubleV2(p)
	fmt.Println(*p)

	var sl []int = []int{1, 2, 3}
	DoubleV3(sl)
	fmt.Println(sl) // 参照型のデータ構造の値は上書きされる
}
