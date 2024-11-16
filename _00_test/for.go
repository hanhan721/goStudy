package main

import (
	"fmt"
)

// 简单的平方根方法，只求出整数位
func Sqrt(x float64) float64 {
	z := 1.0
	for {
		if z*z > x {
			return z - 1
		}
		z++
	}
}

func main() {
	fmt.Println(Sqrt(10001))
}
