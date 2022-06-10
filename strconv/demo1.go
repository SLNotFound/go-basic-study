package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1)
	}

	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%#v\n", s2, s2)
	b, err := strconv.ParseBool("true")
	f, err := strconv.ParseFloat("3.1415", 64)
	i, err := strconv.ParseInt("21", 2, 64)
	u, err := strconv.ParseUint("2", 10, 64)
	fmt.Println(b)
	fmt.Println(f)
	fmt.Println(i)
	fmt.Println(u)
}
