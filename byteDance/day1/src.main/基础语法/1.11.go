package main

import "fmt"

// 指针
// 拷贝的参数形参
func add3(n int) {
	n += 2
}

// 实参
func add2ptr(n *int) {
	*n += 2
}

func main() {
	n := 5
	add3(n)
	fmt.Println(n) //5
	add2ptr(&n)
	fmt.Println(n) //7
}
