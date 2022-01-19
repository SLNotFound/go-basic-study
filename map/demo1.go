package main

import "fmt"

func main() {
	// cap为8
	m := make(map[string]int, 8)
	m["Alice"] = 21
	m["Bob"] = 22
	fmt.Println(m)
	fmt.Println(m["Alice"])
	fmt.Printf("type of m is: %T\n", m)
}
