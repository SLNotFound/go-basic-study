package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct {
}

func (stu *Student) Speak(msg string) string {
	if msg == "test" {
		return "hello"
	} else {
		return "world"
	}
}

func main() {
	stu := Student{}
	fmt.Println(stu.Speak("rrr"))
}
