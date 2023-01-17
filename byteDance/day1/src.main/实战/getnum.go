package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxMax := 100
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxMax)
	fmt.Println("The secret number is", secretNumber)
}
