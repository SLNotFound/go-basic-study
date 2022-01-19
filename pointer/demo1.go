package main

import "fmt"

func main() {
	var a int
	//打印变量a在内存中的地址
	fmt.Println(&a)

	//定义int类型指针p
	var p *int

	// 将变量a的地址赋给p
	p = &a

	// 对p取值，赋值为20
	*p = 20
	// 变量a的值改为20
	fmt.Println(a)
}
