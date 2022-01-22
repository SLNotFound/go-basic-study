package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine结束就 -1
	fmt.Println("Hello Goroutine! ", i)
}

func main() {
	// 开启10个goroutine，并发执行，每次执行结构都不同。
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就 +1
		go hello(i)
	}

	wg.Wait() // 等待goroutine都结束
}
