package main

import (
	"fmt"
	"time"
)

func task() {
	i := 0
	for {
		i++
		fmt.Println("goroutine执行: i=", i)
		time.Sleep(1 * time.Second)
	}
}

// goroutine 和 main 并行执行代码
func main() {
	//创建一个goroutine 执行task()函数
	go task()

	i := 0
	for {
		i++
		fmt.Println("main执行: i=", i)
		time.Sleep(1 * time.Second)
	}
}
