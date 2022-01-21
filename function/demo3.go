package main

import "fmt"

//接收者为值类型和指针类型的区别

func valueIntTest(a int) int {
	return a + 10
}

func pointerIntTest(a *int) int {
	return *a + 10
}

func structTestValue() {
	a := 2
	fmt.Println("valueIntTest: ", valueIntTest(a))

	b := 5
	fmt.Println("pointerIntTest: ", pointerIntTest(&b))
}

type Person struct {
	id   int
	name string
}

func (p Person) valueShowName() {
	fmt.Println(p.name)
}

func (p *Person) pointerShowName() {
	fmt.Println(p.name)
}

func structTestFunc() {
	personValue := Person{1, "Alice"}
	personValue.valueShowName()
	personValue.pointerShowName()

	personPointer := &Person{102, "Bob"}
	personPointer.valueShowName()
	personPointer.pointerShowName()
}

func main() {
	structTestValue()
	structTestFunc()
}
