package main

import (
	"fmt"
	"time"
)

func getNow() {
	now := time.Now()
	fmt.Println(now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year, int(month), day, hour, minute, second)
}

func getTimestamp() {
	now := time.Now()        // 获取当前时间
	timestamp := now.Unix()  // 秒级时间戳
	milli := now.UnixMilli() // 毫秒时间戳 Go1.17+
	micro := now.UnixMicro() // 微秒时间戳 Go1.17+
	nano := now.UnixNano()   // 纳秒时间戳
	fmt.Println(timestamp)
	fmt.Println(milli)
	fmt.Println(micro)
	fmt.Println(nano)
}

func main() {
	//getNow()
	getTimestamp()
}
