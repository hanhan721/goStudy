package main

import (
	"fmt"
	"sync"
	"time"
)

func test1() {
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

func test2() {
	// 创建一个缓冲区大小为3的通道
	c := make(chan int, 3)
	fmt.Println("通道 c 的当前元素数量:", len(c), "容量:", cap(c))
	go func() {
		defer fmt.Println("go程结束....") //此处defer不会执行,如果给通道的初始化缓冲区>=赋值的个数,就能正常执行defer
		// 往通道塞进去4个值,最后一个会阻塞,等待被取出后才会重新进入通道(ps:队列->先进先出)
		for i := range 4 {
			c <- i
			fmt.Println("给通道c传值:", i)
		}
	}()
	time.Sleep(2 * time.Second) //保证go程能正常执行完
	fmt.Println("通道 c 的当前元素数量:", len(c), "容量:", cap(c))
	for range 3 {
		// 如果通道内没值,还使用<-c进行取值,会报错
		fmt.Println("接收到通道c的值:", <-c)
	}
	// 只从通道取出3个值,最后一个还在通道内
	fmt.Println("通道 c 的当前元素数量:", len(c), "容量:", cap(c))
}

// 测试同时写数据和读数据
var wg sync.WaitGroup

func write(ch chan int) {
	for i := 1; i <= 10; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("写:", i)
		ch <- i
	}
	close(ch)
	wg.Done()
}
func read(ch chan int) {
	// 使用for range 循环读取管道,管道在写入必须手动close, 使用for i 循环就不需要
	for v := range ch {
		// 假设读取速度快于写入速度,管道也没问题
		time.Sleep(100 * time.Millisecond)
		fmt.Println("读:", v)
	}
	wg.Done()
}
func main() {
	//test1()
	//test2()

	//测试同时写数据和读数据
	var ch = make(chan int, 10)
	wg.Add(1)
	go write(ch)
	wg.Add(1)
	go read(ch)
	wg.Wait()
	fmt.Println("程序结束....")
}
