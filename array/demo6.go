package main

import "fmt"

// 二维数组
func main() {
	arr := [...][2]string{
		{"alice", "bob"},
		{"cindy", "dog"},
		{"edison", "frank"},
	}

	fmt.Println(arr)
}
