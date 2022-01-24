package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 以只读方式读取文件
func ReadFile(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file failed, err: ", err)
		return
	}
	fmt.Println("open file successfully!")
	defer f.Close()
}

// 读取文件内容
func ReadFileContent(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file failed, err: ", err)
		return
	}
	defer f.Close()

	var tmp = make([]byte, 128)
	n, err := f.Read(tmp)
	if err == io.EOF {
		fmt.Println("read file failed, err: ", err)
		return
	}

	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(tmp[:n]))
}

// 循环读取文件内容
func ReadFileContentByFor(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file failed, err: ", err)
		return
	}

	defer f.Close()
	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := f.Read(tmp)
		if err == io.EOF {
			fmt.Println("file read eof!!")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err: ", err)
			return
		}
		content = append(content, tmp[:n]...)
		fmt.Printf("共有%d字节\n", n)
	}

	fmt.Println(string(content))

}

func ReadFileByBuffIo(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file failed, err: ", err)
		return
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("file read EOF!!")
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
	filePath := "FileOperation/test.txt"
	//ReadFile(filePath)
	//fmt.Println("============================")
	//ReadFileContent(filePath)
	//fmt.Println("============================")
	//ReadFileContentByFor(filePath)
	//fmt.Println("============================")
	ReadFileByBuffIo(filePath)
}
