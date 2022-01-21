package main

import "fmt"

type Data struct {
	x int
}

func (d Data) ValueTest() {
	fmt.Printf("Value: %p\n", &d)
}

func (d *Data) PointerTest() {
	fmt.Printf("Pointer: %p\n", d)
}

func main() {
	d := Data{}
	p := &d
	fmt.Printf("Data: %p\n", p)

	d.ValueTest()
	d.PointerTest()

	p.ValueTest()
	p.PointerTest()
}
