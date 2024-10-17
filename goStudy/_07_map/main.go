package main

import "fmt"

func main() {
	// 初始化map , key:为int类型  value:string类型
	map1 := make(map[int]string)
	map1[0] = "java"
	map1[1] = "c"
	map1[2] = "python"
	map1[3] = "go"
	fmt.Println(map1)
	// 遍历
	for key, value := range map1 {
		// key是hash结构,无序,类似java的HashMap
		fmt.Println(key, value)
	}
	changeMap(map1) //会修改初始值
	fmt.Println(map1)
}
func changeMap(mmap map[int]string) {
	mmap[1] = "c++"
}
