package main

import "fmt"

// 数组遍历
func main() {
	var arr = [...]string{"alice", "bob", "cindy"}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for index, value := range arr {
		fmt.Printf("index: %d + value: %s\n", index, value)
	}
}
