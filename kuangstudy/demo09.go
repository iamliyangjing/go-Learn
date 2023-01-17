package main

/**
 iota,特殊常量，可以任务是一个可以被编译器修改的常量，
	iota是go语言的常量计数器
  iota在const关键字出现时将被重置为0（const内部的第一行之前），
  const中每新增一行常量声明将使iota技术一次（iota可理解为const语句块中的行索引）
*/

import "fmt"

func main() {
	const (
		a = iota // 0
		b = iota
		c = iota
		d
		e
		f
		g
	)

	const (
		k = iota
		j
	)

	fmt.Println(a, b, c, d, e, f, g)

	fmt.Println(k, j)
}
