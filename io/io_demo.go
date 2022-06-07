package main

import (
	"fmt"
	"io"
	"os"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if err != nil {
		panic(err)
	}
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func main() {
	//reader := strings.NewReader("Hello Golang")
	data, err := ReadFrom(os.Stdin, 12)
	//data, err := ReadFrom(strings.NewReader("Hello Golang"), 12)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
