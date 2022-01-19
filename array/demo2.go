package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 求数组所有元素之和

func sumArr(a [10]int) int {
	var sum = 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func main() {
	// 若想做一个真正的随机数，要种子
	// seed()种子默认是1
	//rand.Seed(1)

	// 获取时间戳
	t := time.Now().Unix()
	fmt.Println(t)
	rand.Seed(t)

	var b [10]int
	for i := 0; i < len(b); i++ {
		// 产生一个0到1000随机数
		b[i] = rand.Intn(1000)
	}
	sum := sumArr(b)
	fmt.Println(b)
	fmt.Printf("sum=%d\n", sum)
}
