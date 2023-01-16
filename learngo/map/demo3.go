package main

import "fmt"

// map的遍历

func main() {
	m := make(map[string]int, 8)
	m["Alice"] = 21
	m["Bob"] = 22
	m["Cindy"] = 23
	for k, v := range m {
		fmt.Printf("k: %s, v:%d\n", k, v)
	}
	fmt.Println("=====================")
	// 只遍历k
	for k := range m {
		fmt.Printf("k: %s\n", k)
	}
	fmt.Println("=====================")
	// 只遍历v
	for _, v := range m {
		fmt.Printf("v: %d\n", v)
	}
}
