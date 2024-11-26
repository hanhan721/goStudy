package main

import (
	"fmt"
	"strconv"
)

func palindromes(a int) bool {
	str := strconv.Itoa(a)
	fmt.Println(str)
	l := len(str)
	for i := range (l + 1) / 2 {
		if str[i] != str[l-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(palindromes(123))
	fmt.Println(palindromes(12321))
	fmt.Println(palindromes(12311))
	fmt.Println(palindromes(1221))
}
