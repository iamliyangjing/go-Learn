package main

// 结构体方法

import "fmt"

type user1 struct {
	name     string
	password string
}

// 移到前面来
func (u user1) checkPassword1(password string) bool {
	return u.password == password
}

// 带指针 可以对结构体进行修改
func (u *user1) resetPassword(password string) {
	u.password = password
}

func main() {
	a := user1{name: "wang", password: "1024"}
	a.resetPassword("2048")
	fmt.Println(a.checkPassword1("1024")) //true
}
