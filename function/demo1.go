package main

import "fmt"

type User struct {
	name string
	age  int
}

func (u *User) Print() {
	fmt.Printf("%v : %v\n", u.name, u.age)
}

func (u User) Print1() {
	fmt.Printf("%v : %v\n", u.name, u.age)
}

func main() {
	u1 := User{
		name: "Alice",
		age:  21,
	}
	u1.Print()

	u2 := User{
		name: "Bob",
		age:  22,
	}
	u2.Print1()
}
