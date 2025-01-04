package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixMilli()
	fmt.Println(find(1, 100000))
	fmt.Println(time.Now().UnixMilli() - start)
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
