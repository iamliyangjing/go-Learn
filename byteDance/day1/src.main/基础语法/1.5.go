package main

import (
	"fmt"
	"time"
)

func main() {
	a := 2
	switch a {
	// 不需要 break
	// 可以加任意变量在expr
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("for")
	default:
		fmt.Println("other")
		t := time.Now()
		switch {
		case t.Hour() < 12:
			fmt.Println("it is before noon")
		default:
			fmt.Println("cao")

		}
	}
}
