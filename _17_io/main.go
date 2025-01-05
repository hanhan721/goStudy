package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	read2()
}

func read() {
	file, err := os.Open("_16_select/main.go")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	//读取文件 一次读取64个字节
	var tempSlice = make([]byte, 128)
	var fileData []byte
	for {
		n, err := file.Read(tempSlice)
		if err == io.EOF {
			// 读取到文件末尾,直接退出循环
			break
		}
		//fileData = append(fileData, tempSlice...)  //这种写法有问题  最后一次读取的字节可能不够定义的128字节,但还是被全部添加到fileData切片中
		fileData = append(fileData, tempSlice[:n]...)
		fmt.Printf("读取到了 %d 个字节\n", n)
	}
	fmt.Println(string(fileData))
}
func read2() {
	file, err := os.Open("_16_select/main.go")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(file)
	readString, _ := reader.ReadString('\n') //使用换行符作为分隔符表示一次读取一行,参考split
	fmt.Println("第一行的数据:", readString)
}
