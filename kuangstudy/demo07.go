package main

import "fmt"

func test() (int, int) {
	return 100, 200
}

// 匿名变量 的特点是 一个下划线"_"，“_”本身就是一个特殊的标志符，
// 被称为空白标识符，它可以像其他标识符那样用户变量的声明或赋值
// 但任何赋值给这个标识符的值都将被抛弃

func main() {
	// 对象：User
	//a, b := test()
	a, _ := test()
	_, b := test()
	fmt.Print(a, b)
}
