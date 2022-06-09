package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func cat(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("read file failed, err: ", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			break
		}
		if err != nil {
			fmt.Println("read file failed, err: ", err)
			return
		}
		fmt.Print(line)
	}
}

func main() {
	fileName := os.Args[1]

	cat(fileName)
}
