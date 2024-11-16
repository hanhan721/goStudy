package main

import "fmt"

//interface空接口

func printItem(item interface{}) {
	fmt.Printf("%s type is %T\n", item, item)
	value, ok := item.(string) //断言语句
	if ok {
		fmt.Println(value, "是字符串类型")
	} else {
		fmt.Println(value == "", "  是其他类型") //value是空字符串
	}
}
func main() {
	printItem(1)
	printItem(3.14)
	printItem("3.14")

}
