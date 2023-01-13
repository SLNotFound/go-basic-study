package main

import "fmt"

// 声明变量,使用var关键字
// 变量名在前，变量类型在后
func declareVariable() {
	var name string
	var age int
	var isExist bool
	fmt.Println(name, age, isExist)
}

// 变量零值
func variableZeroValue() {
	var name string
	var age int
	var isExist bool
	fmt.Printf("name: %q, age: %v, isExist: %v", name, age, isExist)
}

// 短变量声明
func shortVariableDeclare() {
	age := 21
	name := "alice"
	fmt.Printf("%s is %d years old!", name, age)
}

// 常量，关键字const，常量值不变。
func declareConstant() {
	const NAME = "test"
	fmt.Println(NAME)
}

// iota常量计数
func iotaDemo() {
	const (
		a = iota
		b
		c
		d
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func main() {
	//variableZeroValue()
	//shortVariableDeclare()
	//declareConstant()
	iotaDemo()
}
