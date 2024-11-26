package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for range 200 {
		for i := 1; i <= 9; i++ {
			for j := 1; j <= i; j++ {
				fmt.Print(j, "*", i, "=", i*j, "\t")
			}
			fmt.Println()
		}
	}
	fmt.Println("所耗时间:", time.Since(start))
}
