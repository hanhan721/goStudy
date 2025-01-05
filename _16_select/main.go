package main

import (
	"fmt"
	"strconv"
)

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := range 5 {
			//打印c通道中的数据
			fmt.Println("第", i+1, "次写入数据:", <-c)
		}
		//循环5次后给quit写入数据
		quit <- 0
	}()
	selectTest(c, quit)
	fmt.Println("-----------------------")
	intChan := make(chan int, 5)
	stringChan := make(chan string, 5)
	for i := range 5 {
		//往两个管道存入数据
		intChan <- i + 1
		stringChan <- "\"" + strconv.Itoa(i+1) + "\""
	}
	for {
		select {
		case v := <-intChan:
			fmt.Println("从整型管道取出数据:", v)
		case v := <-stringChan:
			fmt.Println("从字符串管道取出数据:", v)
		default:
			fmt.Println("结束....")
			return
		}
	}
}

func selectTest(c, quit chan int) {
	x, y := 1, 2
	for {
		select {
		case c <- x:
			//如果成功给c通道写入数据,则执行case里面的语句
			x = y
			y = x + y
		case <-quit:
			//如果成功给quit通道写入数据
			fmt.Println("quit")
			return
		}
	}
}
