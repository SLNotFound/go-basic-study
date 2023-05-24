package main

import "fmt"

func main() {
	// 使用长度声明一个字符串切片
	slice1 := make([]string, 5)
	fmt.Println(len(slice1))

	// 使用长度和容量来声明整型切片
	slice2 := make([]int, 3, 5)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	// 通过切片字面量来声明切片
	slice3 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	slice4 := []int{10, 20, 30}
	fmt.Println(slice3)
	fmt.Println(slice4)

	// 使用索引声明切片
	slice5 := []string{99: ""}
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))
}
