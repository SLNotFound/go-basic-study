package main

import "fmt"

func main() {
	// 声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("s1为空")
	} else {
		fmt.Println("s1不为空")
	}

	s2 := []int{}
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)

	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)

	s5 := []int{1, 2, 3}
	fmt.Println(s5)

	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	s6 = arr[1:4]
	fmt.Println(s6)
	fmt.Println(cap(s6))
	fmt.Println(len(s6))
}
