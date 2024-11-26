package main

import "fmt"

func narcissistic(num int) bool {
	if num < 100 || num > 999 {
		return false
	}
	a := num / 100     //百位
	b := num / 10 % 10 //十位
	c := num % 10
	r := a*a*a + b*b*b + c*c*c
	return r == num
}
func main() {
	for i := 100; i < 1000; i++ {
		if narcissistic(i) {
			fmt.Println(i)
		}
	}
}
