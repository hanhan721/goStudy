package main

import "fmt"

func bubbleSort(arr []int) {
	l := len(arr)
	for j := l - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
			}
		}
	}
}
func main() {
	arr := []int{3, 2, 1, 4, 0}
	bubbleSort(arr)
	fmt.Println(arr)
}
