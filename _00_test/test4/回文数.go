package main

import (
	"fmt"
	"strconv"
)

// 转字符串后同时对前后进行比较
func palindromes2(a int) bool {
	str := strconv.Itoa(a)
	fmt.Print(str, " ")
	l := len(str)
	for i := range (l + 1) / 2 {
		if str[i] != str[l-1-i] {
			return false
		}
	}
	return true
}

// 数字反转后比较
func palindromes(num int) bool {
	// 负数和以0结尾的非零数字不是回文
	if num < 0 || (num%10 == 0 && num != 0) {
		return false
	}
	half := 0
	// 截取数字前半部分的反转
	for num > half {
		half = half*10 + num%10
		num /= 10
	}
	return half == num || num == half/10 //奇数的情况,需要再除以10
}

// 反转数字   123 -> 321
func invertNum(num int) int {
	if num < 0 {
		return -1
	}
	var a = 0
	for num > 0 {
		a = a*10 + num%10
		num /= 10
	}
	return a
}

func main() {
	fmt.Println(12304567, palindromes(12304567))
	fmt.Println(10, palindromes(10))
	fmt.Println(123, palindromes(123))
	fmt.Println(12321, palindromes(12321))
	fmt.Println(12311, palindromes(12311))
	fmt.Println(1221, palindromes(1221))
}
