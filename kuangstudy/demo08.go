package main

import "fmt"

// 常量
// 是一个简单值的标识符，在程序运行时，不会被修改的值
// 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型

func main() {
	const URL string = "www.baidu.com"
	const URL2 = "www.baidu.com"

	fmt.Print(URL)
	fmt.Print(URL2)
}