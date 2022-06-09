package main

import (
	"log"
	"os"
)

// 创建logger
func main() {
	logger := log.New(os.Stdout, "<INFO>\t", log.LstdFlags)
	logger.Println("这是自定义的logger记录的日志")
}
