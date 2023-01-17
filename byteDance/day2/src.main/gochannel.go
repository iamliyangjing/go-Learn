package main

func main() {
	src := make(chan int)
	dest := make(chan int, 3)
	// A
	go func() {
		defer close(src)

		for i := 0; i < 10; i++ {
			src <- i
		}
	}()
	// B
	go func() {
		defer close(dest)
		for i := range src {
			dest <- i * i
		}
	}()

	// M
	for i := range dest {
		//复杂操作
		println(i)
	}
}
