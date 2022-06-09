package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeToFile1() {
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err: ", err)
		return
	}
	defer file.Close()
	str := "superman\n"
	file.Write([]byte(str))
	file.WriteString("flash")
}

func writeToFile2() {
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err: ", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello world\n") // 将数据写入缓存
	}
	writer.Flush() // 将缓存中的数据写入文件
}

func writeToFile3() {
	str := "hello world"
	err := ioutil.WriteFile("./test.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func main() {
	writeToFile3()
}
