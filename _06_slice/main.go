package main

import "fmt"

func main() {
	array1 := [4]int{}
	for i := 0; i < len(array1); i++ {
		array1[i] = i
	}
	//方法内修改数组,实际数组并没变化
	changeArray1(array1)
	for index, value := range array1 {
		fmt.Println("index:", index, "  value:", value)
	}

	fmt.Println("---------------")
	array := []int{1, 2, 3, 4} //动态数组,不指定长度,切片slice
	fmt.Printf("array type is %T\n", array)
	//方法内修改数组,实际数组会变化
	changeArray(array)
	for _, value := range array {
		fmt.Println(value)
	}

	fmt.Println("------------")

	//var slice []int
	//slice = make([]int, 3)
	slice := make([]int, 3) //初始化长度3(长度3)
	//slice[4]=4 //会报错
	slice[0] = 0
	slice[1] = 1
	slice[2] = 2
	slice = append(slice, 3) //使用这个方法新增元素
	fmt.Println(slice)

	//深拷贝
	var slice2 = make([]int, 10)
	copy(slice2, slice)
	fmt.Println(slice2)
}

// 值传递
func changeArray1(array [4]int) {
	array[0] = 111
	//方法内已经修改
	fmt.Println(array)
}

func changeArray(array []int) {
	array[0] = 111
	//方法内已经修改
	fmt.Println(array)
}
