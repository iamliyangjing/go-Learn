package main

import "fmt"

func main() {
	//var 定义变量 ，如果没有赋值 变量就默认值
	var (
		name string
		age  int
		addr string
	)
	//string 默认值：空
	// int 默认值： 0

	name = "liyangjing"
	age = 25
	addr = "中国"

	/*
		1. 整形和浮点型变量的默认值为0和0.0
		2. 字符串变量的默认值为空字符串
		3. 布尔型变量默认为false
		4. 切片、函数、指针变量的默认为nil
	*/
	fmt.Println(name, addr, age)
}
