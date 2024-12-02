package main

import "fmt"

func binarySearch(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		t := (i + j) / 2
		if nums[t] == target {
			return t
		} else if nums[t] > target {
			j = t - 1
		} else {
			i = t + 1
		}
	}
	return -1
}
func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(binarySearch(nums, 2))
	fmt.Println(binarySearch(nums, 5))
	fmt.Println(binarySearch(nums, 9))
	fmt.Println(binarySearch(nums, 25))
}
