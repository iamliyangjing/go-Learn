package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	t := time.Date(2022, 3, 27, 1, 25, 36, 0, time.UTC)
	t2 := time.Date(2022, 3, 27, 2, 30, 36, 0, time.UTC)
	fmt.Println(t)
	// 时间日期
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour())
	// 时间格式化
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	// 两个时间段的时间间隔
	diff := t2.Sub(t)
	fmt.Println(diff)                           // 1h5m3s
	fmt.Println(diff.Minutes(), diff.Seconds()) //6s 3900
	t3, err := time.Parse("2006-01-02 15:04:05", "2022-03-27 15:04:05")
	if err != nil {
		panic(err)
	}
	fmt.Println(t3 == t)    //true
	fmt.Println(now.Unix()) // 时间戳
}
