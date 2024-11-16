package main

import "fmt"

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
