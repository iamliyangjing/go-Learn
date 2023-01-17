package main

import "fmt"

func main() {
	// := 自动推导
	// 短变量声明并初始化
	name2 := "kuangshen"
	age2 := 18
	fmt.Println("name : %s,age: %d", name2, age2)
	//查看类型
	fmt.Printf("%T,%T", name2, age2)
}
