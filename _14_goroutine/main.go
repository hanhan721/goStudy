package main

import (
	"fmt"
	"runtime"
	"sync"
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

var wg sync.WaitGroup //定义WaitGroup变量

// goroutine 和 main 并行执行代码
func main() {
	/*	//创建一个goroutine 执行task()函数
		go task()
		time.Sleep(500 * time.Millisecond)

		i := 0
		for {
			i++
			fmt.Println("main执行: i=", i)
			time.Sleep(1 * time.Second)
		}*/
	numCPU := runtime.NumCPU()
	fmt.Println("CPU总核数", numCPU)
	for i := 1; i <= 20; i++ {
		wg.Add(1) //新增一个协程
		go test(i)
	}
	wg.Wait()
	fmt.Println("结束.....")
}

func test(num int) {
	for i := range 3 {
		fmt.Printf("协程 %v 打印的第 %v 条数据\n", num, i)
		time.Sleep(100 * time.Millisecond)
	}
	defer wg.Done() //执行完协程减1
}
