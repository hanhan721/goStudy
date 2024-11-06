package main

import (
	"fmt"
	"time"
)

func main() {
	//创建通道
	c := make(chan int)

	//创建一个go程
	go func(a int, b int) {
		fmt.Println("goroutine 执行...")
		time.Sleep(3 * time.Second) //等待3秒
		c <- a + b                  //将a+b的值通过管道发送给c (main程)
	}(100, 50)

	start := time.Now()
	num := <-c //阻塞等待go程中发送数据到c中
	end := time.Now()
	fmt.Println("阻塞时长:", end.Second()-start.Second(), "秒")
	fmt.Println("num=", num)
}
