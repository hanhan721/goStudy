package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(returns(1, 2))
	fmt.Println(returns2(1, 2))
}

// 两数和方法
func sum(a int, b int) int {
	return a + b
}

// 多返回,匿名
func returns(a int, b int) (string, string) {
	return strconv.Itoa(a), strconv.Itoa(b) //int转字符串返回
}

// 多返回,有形参名称
func returns2(a int, b int) (r1 int, r2 int) {
	r1 = a * 10
	r2 = r1 * b
	return
}
