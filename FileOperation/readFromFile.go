package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFromFile1() {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file failed, err: ", err)
		return
	}
	defer file.Close()

	// Read方法来读取文件内容
	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读取完毕")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err: ", err)
			return
		}
		content = append(content, tmp[:n]...)
	}

	fmt.Println(string(content))
}

func readFromFile2() {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file failed, err: ", err)
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
			fmt.Println("文件读取完毕！")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err: ", err)
			return
		}
		fmt.Print(line)
	}
}

func readFromFile3() {
	content, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Println("read file failed, err: ", err)
		return
	}
	fmt.Println(string(content))
}

func main() {
	readFromFile3()
}
