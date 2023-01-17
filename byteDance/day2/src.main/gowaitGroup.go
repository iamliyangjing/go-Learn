package main

import (
	"fmt"
	"sync"
)

func hello1(i int) {
	println("hello goroutine: " + fmt.Sprint(i))
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			defer wg.Done()
			hello1(j)
		}(i)
	}
	wg.Wait()
	println("运行完成")
}
