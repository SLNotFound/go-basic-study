package main

import (
	"fmt"
	"log"
	"os"
)

// 配置日志输出位置
func main() {
	logFile, err := os.OpenFile("./test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed: err: ", err)
		return
	}

	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags)
	log.Println("这是一条普通的日志")
	log.SetPrefix("[INFO]\t")
	log.Println("这是一条普通的日志")
}
