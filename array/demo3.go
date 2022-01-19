package main

import "fmt"

// 练习 给出一个数组，找出两个元素相加为给定数值的下标

func findArr(arr [6]int, target int) {

	// 遍历arr
	for i := 0; i < len(arr); i++ {
		// 用给定的数值减去 arr遍历的元素
		other := target - arr[i]
		//从后一个元素遍历，匹配差值
		for j := i + 1; j < len(arr); j++ {
			if arr[j] == other {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}

func main() {
	a := [6]int{2, 3, 5, 6, 8, 0}
	findArr(a, 8)
}
