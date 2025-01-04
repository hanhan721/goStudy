package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now().UnixMicro()
	fmt.Println(find(1, 100000))
	end := time.Now().UnixMicro()
	fmt.Println(end - start)
	//使用管道进行查找素数
	ch := make(chan int, 1000)
	var primeNums []int
	wg.Add(1)
	go putNums(ch, 100000)
	wg.Add(5)
	primeMutex := &sync.Mutex{} // 保护 primeNums 的互斥锁  -> 如果使用管道存储素数就不需要加锁
	for range 5 {
		// 使用5个协程消费管道,速度比不单协程计算快
		go getNums(ch, &primeNums, primeMutex)
	}
	wg.Wait()
	fmt.Println(primeNums)
	fmt.Println(time.Now().UnixMicro() - end)

}

// 找素数
func find(start, end int) []int {
	var nums []int
tags:
	for i := start; i <= end; i++ {
		if i < 2 { // 1 不是素数
			continue
		}
		for j := 2; j < i; j++ {
			if i%j == 0 {
				continue tags
			}
		}
		nums = append(nums, i)
	}
	return nums
}

// 使用管道
func putNums(ch chan int, n int) {
	for i := 1; i <= n; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}
func getNums(ch chan int, primeNums *[]int, mutex *sync.Mutex) {
	for v := range ch {
		if isPrimeNumber(v) {
			mutex.Lock()
			*primeNums = append(*primeNums, v)
			mutex.Unlock()
		}
	}
	wg.Done()
}
func isPrimeNumber(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
