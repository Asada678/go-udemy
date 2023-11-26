package main

import (
	"fmt"
	. "fmt"
	f "fmt"

	"lesson12/foo"
)

// scope

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	// var b string = s
	b = s
	{
		b := "BBBB"
		fmt.Println(b)
	}
	fmt.Println(b)
	return b
}

func main() {
	fmt.Println(foo.Max)
	// f.Println(foo.min) // 定義された変数が小文字から始まる場合、外部から参照できない

	f.Println(foo.ReturnMin()) // 関数が呼び出しているパッケージ内の小文字変数は参照できる
	Println(foo.Max)

	fmt.Println(appName())
	// fmt.Println(AppName, Version)]

	fmt.Println(Do("AAA"))
}
