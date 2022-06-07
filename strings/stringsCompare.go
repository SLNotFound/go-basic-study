package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "Hello"
	s2 := "Bob"
	fmt.Println(strings.Compare(s1, s2))
	fmt.Println(strings.Compare(s1, s1))
	fmt.Println(strings.Compare(s2, s1))

	fmt.Println(strings.EqualFold("GO", "go"))
	fmt.Println(strings.EqualFold("一", "二"))

}
