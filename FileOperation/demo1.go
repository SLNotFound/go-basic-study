package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed")
		return
	}
	defer file.Close()
	fmt.Println("open file successfully")
}
