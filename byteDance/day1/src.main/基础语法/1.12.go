package main

import "fmt"

type user struct {
	name     string
	password string
}

func main() {
	a := user{name: "wang", password: "1024"}
	b := user{"wang", "1024"}
	c := user{name: "Wange"}
	c.password = "1024"

	var d user
	d.password = "1024"
	d.name = "li"

	fmt.Println(a, b, c, d)
	fmt.Println(checkPassword(a, "haha"))
	fmt.Println(checkpassword2(&a, "haha"))
}

func checkPassword(u user, password string) bool {
	return u.password == password
}

func checkpassword2(u *user, password string) bool {
	return u.password == password
}
