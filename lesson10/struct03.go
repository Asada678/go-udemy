package main

import "fmt"

type T struct {
	// User User
	User // 型名を省略できる
}

type User struct {
	Name string
	Age  int
}

func (u *User) SetName() {
	u.Name = "A"
}

func main() {
	t := T{User: User{Name: "user1", Age: 10}}

	fmt.Println(t)
	fmt.Println(t.User)
	fmt.Println(t.User.Name)

	fmt.Println(t.Name) // 型名を省略した場合、Userを省略してNameにアクセスできる

	t.User.SetName()
	fmt.Println(t)

}
