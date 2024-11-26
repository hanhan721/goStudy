package main

import (
	"fmt"
	"math"
	"time"
)

// 获取最小公倍数
func getLeastCommonMultiple(a, b int) int {
	return a * b / getGreatestCommonDivisor(a, b)
}

// 获取最大公约数
func getGreatestCommonDivisor(a, b int) int {
	a = int(math.Abs(float64(a)))
	b = int(math.Abs(float64(b)))

	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

func getGreatestCommonDivisor2(a, b int) int {
	//性能低
	r := min(a, b)
	if r < 1 {
		return -1
	}
	for i := r; i > 1; r-- {
		if a%r == 0 && b%r == 0 {
			return r
		}
	}
	return 1
}

func main() {
	a, b := 128128128, 96969696
	start := time.Now()
	r := getGreatestCommonDivisor(a, b)
	fmt.Println("所耗时间:", time.Since(start))
	start = time.Now()
	r2 := getGreatestCommonDivisor2(a, b)
	fmt.Println("所耗时间:", time.Since(start))
	fmt.Println(a, "和", b, "的最大公约数是:", r, r2)

	r3 := getLeastCommonMultiple(100, 200)
	r4 := getGreatestCommonDivisor2(100, 200)
	fmt.Println("最大公约数是", r4, "最小公倍数是:", r3)
}
