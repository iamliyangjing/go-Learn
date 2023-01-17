package main

import (
	"sync"
	"time"
)

var (
	x    int64
	lock sync.Mutex
)

func addWithLock() {
	for i := 0; i < 2000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}
func addWriteoutLock() {
	for i := 0; i < 2000; i++ {
		x += 1
	}
}

// 测试
func main() {
	x = 0
	// 没有加锁出现并发安全问题
	for i := 0; i < 5; i++ {
		go addWriteoutLock()
	}

	time.Sleep(time.Second)
	println("WithOutLock:", x)
	x = 0
	for i := 0; i < 5; i++ {
		go addWithLock()
	}

	time.Sleep(time.Second)
	println("WithLock:", x)
}
