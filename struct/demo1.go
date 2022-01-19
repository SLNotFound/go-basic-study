package main

import "fmt"

type student struct {
	name string
	age  int
	city string
}

func main() {
	var stu1 student
	stu1.name = "Alice"
	stu1.age = 21
	stu1.city = "Shanghai"
	fmt.Printf("stu1= %v\n", stu1)
	// 输出student结构体里的变量
	fmt.Printf("stu1= %#v\n", stu1)
	fmt.Println("===========================")
	//匿名结构体
	var user struct {
		name string
		age  int
	}
	user.name = "Bob"
	user.age = 22
	fmt.Printf("user : %#v\n", user)
	fmt.Println("===========================")

	// 指针类型结构体
	var stu2 = new(student)
	stu2.name = "Cindy"
	stu2.age = 23
	stu2.city = "Nanjing"
	fmt.Printf("%T\n", stu2)
	fmt.Printf("stu2= %#v\n", stu2)
	fmt.Println("===========================")

	// 结构体的地址实例化
	stu3 := &student{}
	fmt.Printf("%T\n", stu3)
	fmt.Printf("stu3 = %#v\n", stu3)
	stu3.name = "Dog"
	stu3.age = 24
	stu3.city = "Beijing"
	fmt.Printf("stu3 = %#v\n", stu3)
}
