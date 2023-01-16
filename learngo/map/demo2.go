package main

import "fmt"

// 判断map中键是否存在

func main() {
	m := make(map[string]int, 8)
	m["Alice"] = 21
	m["Bob"] = 22

	v, ok := m["Alice"]
	//v, ok := m["Cindy"] // 返回 no no no
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("no no no")
	}
}
