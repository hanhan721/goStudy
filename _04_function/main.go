package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(returns(1, 2))
	fmt.Println(returns2(1, 2))
	fmt.Println(sums(1, 2, 3, 4, 5))
	fmt.Println("_______________")
	var func1 callable
	func1 = do(1)
	fmt.Println(func1(3, 4))
	func1 = do(2)
	fmt.Println(func1(3, 4))
	fmt.Println("________________")
	recursionPrint(10)
	fmt.Println()
	fmt.Println(recursionSum(5))

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

// 不确定参数
func sums(x ...int) int {
	sum := 0
	for i := range x {
		sum += i
	}
	return sum
}

// 定义直接返回一个方法
type callable func(int, int) int

// func do(x int) func(int, int) int {
func do(x int) callable {
	switch x {
	case 1: //可以是已经定义的方法
		return sum
	case 2:
		//也可以是匿名方法
		return func(a, b int) int {
			return a - b
		}
	default:
		return sum
	}
}

// 递归打印0到n的所有整数
func recursionPrint(n int) {
	if n > 0 {
		n -= 1
		recursionPrint(n)
		fmt.Print(n, "\t")
	}
}

// 递归获取0到n的和
func recursionSum(n int) int {
	if n > 0 {
		return n + recursionSum(n-1)
	} else {
		return 0
	}
}
