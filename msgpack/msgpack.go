package main

import (
	"fmt"
	"github.com/vmihailenco/msgpack"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	p1 := Person{
		Name:   "alice",
		Age:    18,
		Gender: "男",
	}
	// marshal
	b, err := msgpack.Marshal(p1)
	if err != nil {
		fmt.Printf("msgpack marshal failed, err:%v", err)
		return
	}

	// unmarshal
	var p2 Person
	err = msgpack.Unmarshal(b, &p2)
	if err != nil {
		fmt.Printf("msgpack unmarshal failed, err:%v", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}
