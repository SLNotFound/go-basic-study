package main

import "fmt"

type TestMethod interface {
	say()
	run()
}

type dog struct {
	name string
}

type cat struct {
	name string
}

func (d dog) say() {
	fmt.Printf("dog dog dog")
}

func (d dog) run() {
	fmt.Printf("run1 run1 run1")
}

func (c cat) say() {
	fmt.Printf("cat cat cat")
}

func (c cat) run() {
	fmt.Printf("run2 run2 run2")
}

func main() {
	d := dog{"alice"}
	d.say()
	d.run()
	fmt.Println("\n===============")
	c := cat{"bob"}
	c.say()
	c.run()
}
