package main

import (
	"errors"
	"fmt"
)

type user3 struct {
	name     string
	password string
}

func findUser(user3s []user3, name string) (v *user3, err error) {
	for _, u := range user3s {
		if u.name == name {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}

func main() {
	u, err := findUser([]user3{{"wang", "1024"}}, "wang")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.name) // wang

	//if u,err :=findUser([]user3{{"wang","1024"}},"li");
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}else {
	//	fmt.Println(u.name)
	//}
}
