package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		sub := target - num
		if _, ok := m[sub]; ok {
			return []int{i, m[sub]}
		}
		m[num] = i
	}
	return []int{-1}
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5}, 7))
}
