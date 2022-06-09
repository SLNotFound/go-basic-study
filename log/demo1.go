package main

import "log"

func main() {
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("这是一条很普通的日志")
}
