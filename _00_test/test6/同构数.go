package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 求同构数1
func isomorphic2(num int) bool {
	n := strconv.Itoa(num * num)
	numStr := strconv.Itoa(num)
	i := strings.LastIndex(n, numStr)

	return len(n)-len(numStr) == i
}

// 求同构数2
func isomorphic(num int) bool {
	square := num * num
	// 计算 num 的位数
	numDigits := getDigits(num)
	// 取 square 的最后 numDigits 位
	mask := intPow(10, numDigits)
	return square%mask == num
}

// 计算 10 的 n 次方
func intPow(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

// 获取数字位数
func getDigits(num int) int {
	i := 0
	for num != 0 {
		num = num / 10
		i++
	}
	return i
}

func main() {
	for i := 1; i < 1000; i++ {
		if isomorphic(i) {
			fmt.Println(i)
		}
	}
}
