package main

import "log"

// 配置日志的前缀

func main() {
	log.SetFlags(log.LstdFlags)
	log.Println("这是一条普通的日志")
	log.SetPrefix("[INFO]\t")
	log.Println("这是一条普通的日志")
}
